package services

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	//"github.com/beego/beego/v2/adapter/logs"
	"os"
	"reflect"
)

func GetTwitterClient() (client *twitter.Client) {
	// 大きい struct, array は値をコピーするコストが大きくなるからポインタで
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("CONSUMER_KEY"),
		ClientSecret: os.Getenv("CONSUMER_SECRET"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)
	// Twitter client
	client = twitter.NewClient(httpClient)
	return
}

func P(t interface{}) {
	fmt.Println(reflect.TypeOf(t))
}
