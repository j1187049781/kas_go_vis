package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gorm.io/gen"
)

const MySQLDSN = "test:123456@tcp(172.20.1.186:9030)/kgs?charset=utf8mb4&parseTime=True"

func connectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface, // generate mode
	  })
	
	  gormdb := connectDB(MySQLDSN)
	  g.UseDB(gormdb) // reuse your gorm db
	
	  g.ApplyBasic(g.GenerateAllTable()...)
	  
	
	  // Generate the code
	  g.Execute()
}