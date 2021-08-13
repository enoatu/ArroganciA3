package tasks

import (
	"arrogancia/models"
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"strings"
	"time"
)

func OptimizeTweet() {
	o := orm.NewOrm()
	// 重複するツイートを取得
	tweets := []*models.Tweet{}
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("user_id, count(body) as bodyCount").
		From("tweet").
		GroupBy("user_id, body").
		Having("bodyCount > 1").
		ForUpdate()
	sql := qb.String()

	to, err := o.Begin() // TX
	if err != nil {
		logs.Error(err)
		return
	}
	to.Raw(sql).QueryRows(&tweets)
	if len(tweets) == 0 {
		logs.Info("No Deprecated Tweets")
		to.Rollback()
		return
	}

	// 拒否Twitterユーザー追加
	values := []string{}
	now := time.Now()
	deleteUserIds := []int64{}
	insertCount := 0
	for _, v := range tweets {
		deleteUserIds = append(deleteUserIds, v.UserId)
		values = append(values, fmt.Sprintf("(%s,'%s')", strconv.FormatInt(v.UserId, 10), now.Format("2006-01-02 15:04:05")))
		insertCount++
	}
	insertSql := fmt.Sprintf(
		"INSERT INTO deny_twitter_user (user_id, created_on) VALUES %s ON DUPLICATE KEY UPDATE user_id = VALUES(user_id), created_on = VALUES(created_on)",
		strings.Join(values, ", "))
	if _, err = to.Raw(insertSql).Exec(); err != nil {
		logs.Error(err)
		to.Rollback()
		return
	}

	// 重複ツイート削除
	if _, err := to.QueryTable("tweet").Filter("user_id__in", deleteUserIds).Delete(); err != nil {
		logs.Error(err)
		to.Rollback()
		return
	}

	// DELETE FROM user WHERE name = "slene"
	if err = to.Commit(); err != nil {
		logs.Error(err)
		to.Rollback()
		return
	}

	logs.Info("Added %s to candidate denyTwitterUsers", strconv.Itoa(insertCount))
}
