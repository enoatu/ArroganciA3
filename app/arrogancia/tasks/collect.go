package tasks

import (
	"arrogancia/models"
	"arrogancia/services"
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/davecgh/go-spew/spew"
	"github.com/dghubble/go-twitter/twitter"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	Collect()
}

type ngWords []string
type greedWords []string

func Collect() {
	// search tweets
	client := services.GetTwitterClient()
	// q => "検索ワード１　検索ワード２　-exclude:retweets -from:除外するユーザーID from:除外するユーザーID"
	greedWordsQueryStr := greedWords{}.get().getQueryStr()
	query := fmt.Sprintf("%s exclude:retweets exclude:replies filter:safe", greedWordsQueryStr)
	searchTweetParams := &twitter.SearchTweetParams{
		Query:     query,
		TweetMode: "extended",
		Count:     50,
	}

	search, _, err := client.Search.Tweets(searchTweetParams)
	if err != nil {
		logs.Error(err)
		return
	}
	tweets := filterTweets(search.Statuses)
	o := orm.NewOrm()
	for _, v := range tweets {
		// spew.Dump(v)
		createdAt, err := v.CreatedAtTime()
		if err != nil {
			logs.Error(err)
			return
		}
		tweetId, err := strconv.Atoi(v.ID)
		tweetModel := &models.Tweet{
			TweetId:        tweetId,
			SearchWordId:   1,
			Text:           v.FullText,
			UserName:       v.User.Name,
			UserScreenName: v.User.ScreenName,
			CreatedAt:      createdAt,
			CreatedOn:      time.Now(),
		}
		_, err = o.Insert(tweetModel)
		if err != nil {
			logs.Error(err)
			return
		}
	}
	// lastTweet := tweets[len(tweets)-1]
	// fmt.Printf("SEARCH METADATA:\n%+v\n", search.Metadata)
}

func (ns ngWords) get() ngWords {
	// 固定長arrayよりもsliceの方が使いやすい
	ngWordCountStr := os.Getenv("NG_WORDS_COUNT")
	ngWordCount, err := strconv.Atoi(ngWordCountStr)
	if err != nil {
		logs.Warn(err)
		return ns
	}
	i := 2
	for i < ngWordCount {
		ngWordRawStrs := os.Getenv(fmt.Sprintf("NG_WORDS_%s", strconv.Itoa(i)))
		ngWordRaws := strings.Split(ngWordRawStrs, ",")
		ns = append(ns, ngWordRaws...)
		i++
	}
	return ns
}

func (ns ngWords) getQueryStr() (queryStr string) {
	if len(ns) < 2 {
		logs.Warn(ns)
		return
	}
	queryStr += "-" + strings.Join(ns, " -")
	return
}

func (gs greedWords) get() greedWords {
	greedWordRawStrs := os.Getenv("GREED_WORDS")
	gs = strings.Split(greedWordRawStrs, ",")
	return gs
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
