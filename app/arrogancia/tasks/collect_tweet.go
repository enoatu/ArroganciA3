package tasks

import (
	"arrogancia/models"
	"arrogancia/services"
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/orm"
	"github.com/dghubble/go-twitter/twitter"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type ngWords []string
type greedWords []string

const TWEET_SEARCH_COUNT = 100

func Collect() {
	searchWordId := 1
	o := orm.NewOrm()

	lastTweet := &models.Tweet{}
	o.QueryTable("tweet").Filter("SearchWordId", searchWordId).OrderBy("-TweetId").One(lastTweet)

	denyTwitterUsers := []*models.DenyTwitterUser{}
	o.QueryTable("deny_twitter_user").All(&denyTwitterUsers)
	denyTwitterUsersByUserId := make(map[int64]int)
	for _, v := range denyTwitterUsers {
		denyTwitterUsersByUserId[v.UserId] = 1
	}

	client := services.GetTwitterClient()

	// q => "検索ワード１　検索ワード２　-exclude:retweets -from:除外するユーザーID from:除外するユーザーID"
	query := fmt.Sprintf(
		"%s %s exclude:retweets exclude:replies filter:safe",
		"アプリ",
		greedWords{}.get().getQueryStr())

	maxTweetId := int64(0)
	tweetModels := []*models.Tweet{}
	for {
		// 最新のツイート => 古いツイート...の順で取得
		// lastTweetId = 3
		// count = 3
		// 回)
		// 1) 10, 9, 8 の順で取得
		// 2) 7(maxTweetId = 7), 6 5
		// 3) 4(maxTweetId = 4), (3(sinceID = lastTweetId))
		// => 2回目以降、動的にmaxTweetIdを付け替える
		searchTweetParams := &twitter.SearchTweetParams{
			Query:     query,
			TweetMode: "extended",
			SinceID:   lastTweet.TweetId,
			Count:     TWEET_SEARCH_COUNT, // max 100
		}
		if maxTweetId != 0 {
			searchTweetParams.MaxID = maxTweetId - 1 // equal含む
		}

		search, _, err := client.Search.Tweets(searchTweetParams)
		if err != nil {
			logs.Error(err)
			return
		}
		if len(search.Statuses) == 0 {
			// 0の時
			break
		}

		tweets := filterTweets(search.Statuses)
		if len(tweets) == 0 {
			logs.Warn("Not Found Update Tweet In Loop")
			break
		}
		for _, v := range tweets {
			createdAt, err := v.CreatedAtTime()
			if err != nil {
				logs.Error(err)
				return
			}
			if _, ok := denyTwitterUsersByUserId[v.User.ID]; ok {
				// 除外Twitterユーザーならスキップ
				continue
			}
			tweetModel := &models.Tweet{
				TweetId:        v.ID,
				SearchWordId:   searchWordId,
				Body:           v.FullText,
				UserId:         v.User.ID,
				UserName:       v.User.Name,
				UserScreenName: v.User.ScreenName,
				CreatedAt:      createdAt,
				CreatedOn:      time.Now(),
			}
			tweetModels = append(tweetModels, tweetModel)
		}

		gotLastTweetId := search.Statuses[len(search.Statuses)-1].ID
		if len(search.Statuses) < TWEET_SEARCH_COUNT {
			// 抜ける
			break
		}
		maxTweetId = gotLastTweetId
	}
	if len(tweetModels) == 0 {
		return
	}
	_, err := o.InsertMulti(len(tweetModels), tweetModels)
	if err != nil {
		logs.Error(err)
		return
	}
	logs.Info(fmt.Sprintf("%s added New Tweet", strconv.Itoa(len(tweetModels))))
}

func (ns ngWords) get() ngWords {
	// 固定長arrayよりもsliceの方が使いやすい
	ngWordCountStr := os.Getenv("NG_WORDS_COUNT")
	ngWordCount, err := strconv.Atoi(ngWordCountStr)
	if err != nil {
		logs.Warn(err)
		return ns
	}
	for i := 0; i < ngWordCount; i++ {
		ngWordRawStrs := os.Getenv(fmt.Sprintf("NG_WORDS_%s", strconv.Itoa(i)))
		ngWordRaws := strings.Split(ngWordRawStrs, ",")
		ns = append(ns, ngWordRaws...)
	}
	return ns
}

func (ns ngWords) getQueryStr() (queryStr string) {
	queryStr += "-" + strings.Join(ns, " -")
	return
}

func (gs greedWords) get() greedWords {
	greedWordRawStrs := os.Getenv("GREED_WORDS")
	gs = strings.Split(greedWordRawStrs, ",")
	return gs
}

func (gs greedWords) filterRandom() (newGs greedWords) {
	if len(gs) == 0 {
		return gs
	}
	rand.Seed(time.Now().UnixNano()) // rand メソッドの前に実行しないと、Intnの結果が常に1になったりおかしくなる
	num := rand.Intn(len(gs))
	n := make(greedWords, len(gs))
	copy(n, gs)

	for i := 0; i < num; i++ {
		index := rand.Intn(len(n))
		newGs = append(newGs, n[index])
		n = append(n[:index], n[index+1:]...)
	}
	return
}

func (gs greedWords) getQueryStr() (queryStr string) {
	queryStr = strings.Join(gs, " OR ")
	return
}

func filterTweets(tweets []twitter.Tweet) (filteredTweets []twitter.Tweet) {
	ns := ngWords{}.get()
	for _, tweet := range tweets {
		for _, ngWord := range ns {
			if strings.Contains(tweet.FullText, ngWord) {
				break
			}
		}
		filteredTweets = append(filteredTweets, tweet)
	}
	return
}
