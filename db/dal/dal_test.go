package dal_test

import (
	"kas_go_vis/client"
	"kas_go_vis/db/dal"
	"kas_go_vis/db/models"
	"testing"
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
				Address: tag.Address,
				Name:    tag.Name,
				Link:    tag.Link,
			})
		}

		addrBalances = append(addrBalances, models.AddrBalance{
			Address: v.Address,
			Balance: v.Balance,
			Tags:   tags,
		})
	}

	db.Create(&addrBalances)
}