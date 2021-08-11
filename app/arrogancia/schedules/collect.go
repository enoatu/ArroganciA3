package schedules

import (
	"arrogancia/tasks"
	"github.com/astaxie/beego/toolbox"
	"github.com/beego/beego/v2/adapter/logs"
	_ "github.com/go-sql-driver/mysql"
)

func GetCollectTask() *toolbox.Task {
	isRunning := false
	return toolbox.NewTask("collectTask", "0 0 * * * *", func() error {
		if isRunning {
			logs.Warn("running task skip...")
			return nil
		}
		isRunning = true
		logs.Info("CollectTask Start")
		tasks.Collect()
		logs.Info("CollectTask End")
		isRunning = false
		return nil
	})
	// test
	// err := tk.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
