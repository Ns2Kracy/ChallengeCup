package dao

import (
	"log"

	"ChallengeCup/config"
	"ChallengeCup/service/dbmodel"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() *gorm.DB {
	// 读取配置文件
	conf := config.LoadConfig().Mysql
	if conf.Database == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       conf.GetDsn(),
		DefaultStringSize:         256,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}
	// 连接数据库
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		log.Println("mysql connect error: ", err)
		return nil
	}
	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// 同步表结构
	err = db.AutoMigrate(&dbmodel.UserDBModel{})
	if err != nil {
		return nil
	}
	return db
}
