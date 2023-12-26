package main

import (
	"kas_go_vis/client"
	"kas_go_vis/db/dal"
	"kas_go_vis/db/models"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	cron "github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func runGetRichTask(c *client.DefaultClient, db *gorm.DB){
	log.Printf("runGetRichTask")
	richList, err := c.GetRichTopList()
	if err != nil {
		log.Printf("%v",err)
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

	db.Create(&addrBalances)
}

func main() {
	db := dal.ConnectDB(dal.MySQLDSN)
	err := db.AutoMigrate(&models.AddrBalance{}, &models.AddrTag{})
	if err != nil {
		log.Printf("%v",err)
		return
	}

	c := client.NewDefaultClient()

	cronRunTime := cron.New()
	cronRunTime.AddFunc("06 20 * * *", func() { runGetRichTask(c, db) })
	cronRunTime.Start()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}