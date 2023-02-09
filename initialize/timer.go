/**
*@author:yangzhou
*@date: 2023/2/9
*@email: yangzhou2224@shengtian.com
*@description:
 */
package initialize

import (
	"github.com/robfig/cron/v3"
	"yzgin/config"
	"yzgin/global"
)

func Timer() {
	if global.Config.Timer.Start {
		for i := range global.Config.Timer.Detail {
			go func(detail config.Detail) {
				var option []cron.Option
				if global.Config.Timer.WithSeconds {
					option = append(option, cron.WithSeconds())
				}
				//_, err := global.Timer.AddTaskByFunc("ClearDB", global.Config.Timer.Spec, func() {
				//	err := utils.ClearTable(global.GVA_DB, detail.TableName, detail.CompareField, detail.Interval)
				//	if err != nil {
				//		fmt.Println("timer error:", err)
				//	}
				//}, option...)
				//if err != nil {
				//	fmt.Println("add timer error:", err)
				//}
			}(global.Config.Timer.Detail[i])
		}
	}
}
