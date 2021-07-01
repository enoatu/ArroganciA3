package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"os"
	"time"
)

func main() {
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

	tk := toolbox.NewTask("myTask", "* * * * * *", func() error {
		fmt.Println("hello world")
		search, _, _ := client.Search.Tweets(searchTweetParams)
		fmt.Printf("SEARCH TWEETS:\n%+v\n", search.Statuses[0].FullText)
		fmt.Printf("SEARCH METADATA:\n%+v\n", search.Metadata)
		return nil
	})
	err := tk.Run()
	if err != nil {
		fmt.Println(err)
	}
	toolbox.AddTask("myTask", tk)
	toolbox.StartTask()
	time.Sleep(6 * time.Second)
	toolbox.StopTask()
	beego.Run()
}
