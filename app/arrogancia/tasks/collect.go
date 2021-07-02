package tasks

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"os"
)

func main() {
	Collect()
}

func Collect() {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("CONSUMER_KEY"),
		ClientSecret: os.Getenv("CONSUMER_SECRET"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// search tweets
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
