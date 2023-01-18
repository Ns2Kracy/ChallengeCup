package service

import (
	"ChallengeCup/service/dbmodel"

	"gorm.io/gorm"
)

type UserService interface {
	DeleteAll()
	DeleteByUUID(uuid string)
	NewUser(user dbmodel.UserDBModel) dbmodel.UserDBModel
	UpdateUser(user dbmodel.UserDBModel)
	UploadAvatar(uuid string, avatar string)
	IsExistByName(name string) bool
	GetUserByName(name string) dbmodel.UserDBModel
	GetUserByUUID(uuid string) dbmodel.UserDBModel
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		db: db,
	}
}

func (u *userService) DeleteAll() {
	u.db.Delete(&dbmodel.UserDBModel{})
}

func (u *userService) DeleteByUUID(uuid string) {
	u.db.Where("uuid = ?", uuid).Delete(&dbmodel.UserDBModel{})
}

func (u *userService) NewUser(user dbmodel.UserDBModel) dbmodel.UserDBModel {
	u.db.Create(&user)
	return user
}

func (u *userService) UpdateUser(user dbmodel.UserDBModel) {
	u.db.Model(&user).Omit("password").Updates(&user)
}

func (u *userService) UploadAvatar(uuid string, avatar string) {
	u.db.Model(&dbmodel.UserDBModel{}).Where("uuid = ?", uuid).Update("avatar", avatar)
}

func (u *userService) IsExistByName(name string) bool {
	var count int64
	u.db.Model(&dbmodel.UserDBModel{}).Where("username = ?", name).Count(&count)
	return count != 0
}

func (u *userService) GetUserByName(name string) dbmodel.UserDBModel {
	user := dbmodel.UserDBModel{}
	u.db.Where("username = ?", name).First(&user)
	return user
}

func (u *userService) GetUserByUUID(uid string) dbmodel.UserDBModel {
	user := dbmodel.UserDBModel{}
	u.db.Where("uuid = ?", uid).First(&user)
	return user
}
