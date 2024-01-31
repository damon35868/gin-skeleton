package task

import (
	"fmt"
	"gin-skeleton/app/service"

	"github.com/robfig/cron/v3"
)

type OrderTask struct{}

func NewOrderTask() *OrderTask {
	return &OrderTask{}
}

func (orderTask *OrderTask) Run() {
	cronHandler := cron.New(cron.WithSeconds())

	cronHandler.AddFunc("*/30 * * * * *", func() {
		fmt.Println("[Task] 每30秒执行一次")
		service.NewIndexService().Index("Task")
	})
	cronHandler.Start()
}
