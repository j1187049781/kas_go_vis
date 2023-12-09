package dal

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const MySQLDSN = "it:Its123@tcp(172.19.2.151:3307)/kgs?charset=utf8mb4&parseTime=True"

func ConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}

