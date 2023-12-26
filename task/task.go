package task

import (
	"fmt"
	"kas_go_vis/client"
	"kas_go_vis/db/models"
	"time"

	"gorm.io/gorm"
)

type RunRichListTask struct {
	db *gorm.DB
	client client.GrpcClient
	name string
}

func (t *RunRichListTask) Run() {
	fmt.Println("run rich list task")
	// find task and time = today
	// if not exist, create task and time = today
	// if exist, return
	task := models.Task{}
	now := time.Now().Format("2006-01-02")
	t.db.Where("task_name = ? and time = ?", t.name, now).First(&task)
	if task.ID != 0 {
		return
	}

	richList, err := t.client.GetRichTopList()
	if err != nil {
		fmt.Println(err)
		return
	}

	var addrBalances []models.AddrBalance
	for _, v := range richList {
		tags := make([]models.AddrTag, 0)
		for _, tag := range v.Tags {
			tags = append(tags, models.AddrTag{
				Name:    tag.Name,
				Link:    tag.Link,
			})
		}
		
		now := time.Now()

		addrBalances = append(addrBalances, models.AddrBalance{
			Time:  now.Format("2006-01-02 15:04:05"),
			Address: v.Address,
			Balance: v.Balance,
			Tags:   tags,
		})
	}

	t.db.Create(&addrBalances)

	t.db.Create(&models.Task{
		TaskName: t.name,
		Time:     now,
	})
}