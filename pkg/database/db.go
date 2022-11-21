package database

import (
	"fmt"
	"go-fiber-demo/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

func InitDb() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8&collation=utf8mb4_general_ci&parseTime=True&loc=Local&timeout=10000ms"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		SkipInitializeWithVersion: false,
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "tb_",
		},
		QueryFields: true,
	}); err != nil {
		fmt.Println(err.Error())
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetMaxIdleConns(100)
		err := db.AutoMigrate(
			models.Course{},
		)
		if err != nil {
			fmt.Println(err.Error())
		}
		DB = db
	}
}
