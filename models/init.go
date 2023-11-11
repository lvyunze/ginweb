package models

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = ConnectMySQL()
	if err != nil {
		panic(fmt.Errorf("无法连接到MySQL数据库：%s", err))
	}
}

func ConnectMySQL() (*gorm.DB, error) {
	dbHost := viper.GetString("MySQL.Host")
	dbPort := viper.GetString("MySQL.Port")
	dbUser := viper.GetString("MySQL.Username")
	dbPassword := viper.GetString("MySQL.Password")
	dbName := viper.GetString("MySQL.DbName")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("无法连接到MySQL数据库：%s", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("无法获取数据库连接：%s", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db, nil
}
