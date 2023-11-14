package models

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dbHost := viper.GetString("MySQL.Host")
	dbPort := viper.GetString("MySQL.Port")
	dbUser := viper.GetString("MySQL.Username")
	dbPassword := viper.GetString("MySQL.Password")
	dbName := viper.GetString("MySQL.DbName")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("无法连接到MySQL数据库：%v", err)
		panic("无法连接到MySQL数据库")
	}

	// 设置连接池参数
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Printf("获取数据库连接池失败：%v", err)
		panic("获取数据库连接池失败")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}
