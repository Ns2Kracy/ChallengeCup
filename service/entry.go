package service

import (
	"ChallengeCup/dao"

	"gorm.io/gorm"
)

type ServiceEntry struct {
	UserService UserService
	MqttService MqttService
}

func NewServiceEntry(db *gorm.DB) *ServiceEntry {
	return &ServiceEntry{
		UserService: NewUserService(db),
		MqttService: NewMqttService(db),
	}
}

var Service = NewServiceEntry(dao.DB)
