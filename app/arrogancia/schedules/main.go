package schedules

import (
	"arrogancia/tasks"
	"github.com/astaxie/beego/toolbox"
	"github.com/beego/beego/v2/adapter/logs"
	_ "github.com/go-sql-driver/mysql"
)

func Start() {
	toolbox.AddTask("collectTweetTask", GetCollectTweetTask())
	toolbox.AddTask("optimizeTweetTask", GetOptimizeTweetTask())
	toolbox.StartTask()
}

func GetCollectTweetTask() *toolbox.Task {
	isRunning := false
	return toolbox.NewTask("collectTweetTask", "0 * * * * *", func() error {
		if isRunning {
			logs.Warn("running CollectTweetTask skip...")
			return nil
		}
		isRunning = true
		logs.Info("CollectTweetTask Start")
		tasks.Collect()
		logs.Info("CollectTweetTask End")
		isRunning = false
		return nil
	})
	// test
	// err := tk.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func GetOptimizeTweetTask() *toolbox.Task {
	isRunning := false
	return toolbox.NewTask("optimizeTweetTask", "0 1 * * * *", func() error {
		if isRunning {
			logs.Warn("running OptimizeTweetTask skip...")
			return nil
		}
		isRunning = true
		logs.Info("OptimizeTweetTask Start")
		tasks.OptimizeTweet()
		logs.Info("OptimizeTweetTask End")
		isRunning = false
		return nil
	})
	// test
	// err := tk.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
