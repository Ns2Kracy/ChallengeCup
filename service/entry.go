package service

import "gorm.io/gorm"

var AppService = new(ServiceEntry)

type ServiceEntry struct {
	UserService
}

func NewService(db *gorm.DB) ServiceEntry {
	return ServiceEntry{
		NewUserService(db),
	}
}
