package main

// go run schedules/collect.go だと
// go run: cannot run non-main package なので mainに
import (
	"arrogancia/tasks"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"time"
)

func main() {
	tk := toolbox.NewTask("collectTask", "* * * * * *", func() error {
		fmt.Println("hello world")
		tasks.Collect()
		return nil
	})
	err := tk.Run()
	if err != nil {
		fmt.Println(err)
	}
	toolbox.AddTask("collectTask", tk)
	toolbox.StartTask()
	time.Sleep(6 * time.Second)
	toolbox.StopTask()
	beego.Run()
}
