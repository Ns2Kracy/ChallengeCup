package dao

import (
	log "ChallengeCup/utils/logger"

	"ChallengeCup/config"
	"ChallengeCup/service/dbmodel"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() *gorm.DB {
	conf := config.LoadConfig().Mysql
	if conf.Database == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       conf.GetDsn(),
		DefaultStringSize:         256,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		log.Infof("mysql connect error: ", err)
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	err = db.AutoMigrate(
		&dbmodel.UserDBModel{},
		&dbmodel.MqttData{},
	)
	if err != nil {
		return nil
	}
	log.Info("Mysql connect success")
	return db
}
