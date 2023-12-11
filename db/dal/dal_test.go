package dal_test

import (
	"kas_go_vis/client"
	"kas_go_vis/db/dal"
	"kas_go_vis/db/models"
	"testing"
	"time"
)

func TestDal(t *testing.T) {
	db := dal.ConnectDB(dal.MySQLDSN)
	err := db.AutoMigrate(&models.AddrBalance{}, &models.AddrTag{})
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	c := client.NewDefaultClient()

	richList, err := c.GetRichTopList()
	if err != nil {
		t.Log(err)
		t.Fail()
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