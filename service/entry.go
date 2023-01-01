package service

import (
	"ChallengeCup/dao"

	"gorm.io/gorm"
)

type ServiceEntry struct {
	UserService UserService
}

func NewServiceEntry(db *gorm.DB) *ServiceEntry {
	return &ServiceEntry{
		UserService: NewUserService(db),
	}
}

var Service = NewServiceEntry(dao.InitMysql())