package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type Tweet struct {
	Id             int       `orm:"pk;column(id);auto" json:"id"`
	TweetId        int64     `orm:"column(tweet_id)" description:"ツイートID" json:"tweet_id"`
	SearchWordId   int       `orm:"column(search_word_id)" description:"検索ワードID" json:"search_word_id"`
	Body           string    `orm:"column(body)" description:"ツイート本文" json:"text"`
	UserId         int64     `orm:"column(user_id)" json:"user_id"`
	UserName       string    `orm:"column(user_name)" description:"ユーザーネーム" json:"user_name"`
	UserScreenName string    `orm:"column(user_screen_name)" description:"ユーザースクリーンネーム" json:"user_screen_name"`
	OptimizeLevel  int       `orm:"column(optimize_level)" description:"最適化レベル" json:"optimize_level"`
	CreatedAt      time.Time `orm:"column(created_at);type(datetime)" description:"ツイート作成日時" json:"-"`
	CreatedOn      time.Time `orm:"column(created_on);type(datetime)" description:"作成日時" json:"-"`
	ModifiedOn     time.Time `orm:"column(modified_on);type(timestamp);auto_now_add" description:"最終更新日時" json:"-"`
}

type SearchWord struct {
	Id         int       `orm:"pk;column(id);auto" json:"id"`
	Name       string    `orm:"unique;column(name)" description:"検索ワード" json:"name"`
	CreatedOn  time.Time `orm:"column(created_on);type(datetime)" description:"作成日時" json:"-"`
	ModifiedOn time.Time `orm:"column(modified_on);type(timestamp);auto_now_add" description:"最終更新日時" json:"-"`
}

type DenyTwitterUser struct {
	Id         int       `orm:"pk;column(id);auto" json:"id"`
	UserId     int64     `orm:"unique;column(user_id)" json:"user_id"`
	CreatedOn  time.Time `orm:"column(created_on);type(datetime)" description:"作成日時" json:"-"`
	ModifiedOn time.Time `orm:"column(modified_on);type(timestamp);auto_now_add" description:"最終更新日時" json:"-"`
}

func init() {
	orm.RegisterModel(
		new(Tweet),
		new(SearchWord),
		new(DenyTwitterUser))
}
