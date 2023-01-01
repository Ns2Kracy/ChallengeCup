package service

import (
	"ChallengeCup/service/dbmodel"

	"gorm.io/gorm"
)

type UserService interface {
	DeleteAll()
	DeleteById(id string)
	NewUser(user dbmodel.UserDBModel) dbmodel.UserDBModel
	UpdateUser(user dbmodel.UserDBModel)
	UploadAvatar(id string, avatar string)
	IsExistByName(name string) bool
	GetUserByName(name string) dbmodel.UserDBModel
	GetUserById(id string) dbmodel.UserDBModel
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

func (u *userService) DeleteById(id string) {
	u.db.Where("id = ?", id).Delete(&dbmodel.UserDBModel{})
}

func (u *userService) NewUser(user dbmodel.UserDBModel) dbmodel.UserDBModel {
	u.db.Create(&user)
	return user
}

func (u *userService) UpdateUser(user dbmodel.UserDBModel) {
	u.db.Model(&user).Omit("password").Updates(&user)
}

func (u *userService) UploadAvatar(id string, avatar string) {
	u.db.Model(&dbmodel.UserDBModel{}).Where("id = ?", id).Update("avatar", avatar)
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

func (u *userService) GetUserById(id string) dbmodel.UserDBModel {
	user := dbmodel.UserDBModel{}
	u.db.Where("id = ?", id).First(&user)
	return user
}
