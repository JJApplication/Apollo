/*
Project: Apollo task_test.go
Created: 2021/12/28 by Landers
*/

package task_manager

import (
	"fmt"
	"testing"
	"time"

	"github.com/JJApplication/Apollo/logger"
)

func init() {
	logger.InitLogger()
}

func TestTask_Start(t *testing.T) {
	task := NewTask("test", 0)
	task.Fn = func() {
		fmt.Println("no timeout")
	}
	uuid, err := task.Start()
	fmt.Println(uuid, err)

	task1 := NewTask("test1", 1)
	task1.Fn = func() {
		fmt.Println("with timeout")
		time.Sleep(2 * time.Second)
	}
	uuid, err = task1.Start()
	time.Sleep(5 * time.Second)
	fmt.Println(uuid, err)

	taskInfo, err := task.Info()
	t.Logf("%+v %v", taskInfo, err)
	taskInfo1, err := task1.Info()
	t.Logf("%+v %v", taskInfo1, err)
}
