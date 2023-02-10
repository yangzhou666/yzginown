/**
*@author:yangzhou
*@date: 2022/12/11
*@email: yangzhou2224@shengtian.com
*@description:
 */
package schedule

import (
	"context"
	"fmt"
	"testing"
)

func TestCron(t *testing.T) {
	c := NewServer()

	c.c.AddFunc("@every 1s", func() {
		fmt.Println("我是定时任务")
	})

	c.Start(context.Background())

	select {}
}
