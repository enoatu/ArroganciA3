package tasks

import (
	"arrogancia/services"
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/dghubble/go-twitter/twitter"
	"os"
	// "net/url"
	// "unicode/utf8"
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
	// API仕様としてクエリは500文字以内?
	// if length := utf8.RuneCountInString(url.QueryEscape(query)); length > 1000 {
	// 	logs.Warn(fmt.Sprintf("query string length(%s) is over 1000", strconv.Itoa(length)))
	// 	return
	// }
	searchTweetParams := &twitter.SearchTweetParams{
		Query:     query,
		TweetMode: "extended",
		Count:     3,
	}

	search, _, err := client.Search.Tweets(searchTweetParams)
	if err != nil {
		logs.Warn(err)
		return
	}
	services.P(search.Statuses)
	for _, v := range search.Statuses {
		fmt.Printf("SEARCH TWEETS:\n%+v\n", v.FullText)
	}
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

// func (statues search.Statuses) filterWords (search.Statuses) {
// 	ngWordCountStr := os.Getenv("EXCEPT_WORDS_COUNT")
// 	ngWordCount, err := strconv.Atoi(ngWordCountStr)
// 	if err != nil {
// 		logs.Warn(err)
// 		return ns
// 	}
// 	i := 0
// 	for i < ngWordCount {
// 		ngWordRawStrs := os.Getenv(fmt.Sprintf("NG_WORDS_%s", strconv.Itoa(i)))
// 		logs.Warn(ngWordRawStrs)
// 		logs.Warn(fmt.Sprintf("NG_WORDS_%s", i))
// 		ngWordRaws := strings.Split(ngWordRawStrs, ",")
// 		ns = append(ns, ngWordRaws...)
// 		i++
// 	}
// 	return ns
// }
