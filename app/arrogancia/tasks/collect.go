package tasks

import (
	"arrogancia/services"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
)

func main() {
	Collect()
}

func Collect() {
	// search tweets
	client := services.GetTwitterClient()
	// q => '検索ワード１　検索ワード２　-exclude:retweets -from:除外するユーザーID from:除外するユーザーID'
	searchTweetParams := &twitter.SearchTweetParams{
		Query:     "眠い",
		TweetMode: "extended",
		Count:     3,
	}

	fmt.Println("hello world")
	search, _, _ := client.Search.Tweets(searchTweetParams)
	fmt.Printf("SEARCH TWEETS:\n%+v\n", search.Statuses[0].FullText)
	fmt.Printf("SEARCH METADATA:\n%+v\n", search.Metadata)
}
