package tasks

import (
	"arrogancia/services"
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/davecgh/go-spew/spew"
	"github.com/dghubble/go-twitter/twitter"
	"os"
	"strconv"
	"strings"
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
	logs.Warn(query)
	searchTweetParams := &twitter.SearchTweetParams{
		Query:     query,
		TweetMode: "extended",
		Count:     50,
	}

	search, _, err := client.Search.Tweets(searchTweetParams)
	if err != nil {
		logs.Warn(err)
		return
	}
	services.P(search.Statuses)
	tweets := filterTweets(search.Statuses)
	for _, v := range tweets {
		spew.Dump(v)
	}
	lastTweet := tweets[len(tweets)-1]
	logs.Info(lastTweet)
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
