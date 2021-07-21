package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type Tweet struct {
	Id             int       `orm:"pk;column(id);auto" json:"id"`
	TweetId        int       `orm:"unique;column(tweet_id)" description:"ツイートID" json:"tweet_id"`
	SearchWordId   int       `orm:"column(search_word_id)" description:"検索ワードID" json:"search_word_id"`
	Text           string    `orm:"column(text)" description:"ツイート本文" json:"text"`
	UserName       string    `orm:"column(user_name)" description:"ユーザーネーム" json:"user_name"`
	UserScreenName string    `orm:"column(user_screen_name)" description:"ユーザースクリーンネーム" json:"user_screen_name"`
	CreatedAt      time.Time `orm:"column(created_at);type(datetime)" description:"ツイート作成日時" json:"-"`
	CreatedOn      time.Time `orm:"column(created_on);type(datetime)" description:"作成日時" json:"-"`
	ModifiedOn     time.Time `orm:"column(modified_on);type(timestamp);auto_now_add" description:"最終更新日時" json:"-"`
}

func (t *Tweet) TableIndex() [][]string {
	return [][]string{
		[]string{"SearchWordId"},
	}
}

type SearchWord struct {
	Id         int       `orm:"pk;column(id);auto" json:"id"`
	Word       string    `orm:"unique;column(word)" description:"検索ワード" json:"word"`
	CreatedOn  time.Time `orm:"column(created_on);type(datetime)" description:"作成日時" json:"-"`
	ModifiedOn time.Time `orm:"column(modified_on);type(timestamp);auto_now_add" description:"最終更新日時" json:"-"`
}

func init() {
	orm.RegisterModel(new(Tweet), new(SearchWord))
}
