package service

import (
	"ChallengeCup/service/dbmodel"

	"gorm.io/gorm"
)

type UserService interface {
	DeleteByUUID(uuid string)
	NewUser(user dbmodel.UserDBModel) dbmodel.UserDBModel
	UpdateUser(user dbmodel.UserDBModel)
	UpdatePassword(uuid string, password string)
	UploadAvatar(uuid string, avatar string)
	IsExistByName(name string) bool
	IsExistByUUID(uuid string) bool
	IsExistByPhone(phone string) bool
	IsExistByEmail(email string) bool
	GetUserByName(name string) dbmodel.UserDBModel
	GetUserByUUID(uuid string) dbmodel.UserDBModel
	GetUserByPhone(phone string) dbmodel.UserDBModel
	GetUserByEmail(email string) dbmodel.UserDBModel
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		db: db,
	}
}

// DeleteByUUID 通过UUID删除用户
func (u *userService) DeleteByUUID(uuid string) {
	u.db.Where("uuid = ?", uuid).Delete(&dbmodel.UserDBModel{})
}

// NewUser 创建新用户
func (u *userService) NewUser(user dbmodel.UserDBModel) dbmodel.UserDBModel {
	u.db.Create(&user)
	return user
}

// UpdateUser 更新除了密码以外的用户信息
func (u *userService) UpdateUser(user dbmodel.UserDBModel) {
	u.db.Model(&user).Omit("password").Updates(&user)
}

// UpdatePassword 更新用户密码
func (u *userService) UpdatePassword(uuid string, password string) {
	u.db.Model(&dbmodel.UserDBModel{}).Where("uuid = ?", uuid).Update("password", password)
}

// UploadAvatar 上传用户头像
func (u *userService) UploadAvatar(uuid string, avatar string) {
	u.db.Model(&dbmodel.UserDBModel{}).Where("uuid = ?", uuid).Update("avatar", avatar)
}

// IsExistByName 通过用户名判断用户是否存在
func (u *userService) IsExistByName(name string) bool {
	var count int64
	u.db.Model(&dbmodel.UserDBModel{}).Where("username = ?", name).Count(&count)
	return count != 0
}

// IsExistByUUID 通过用户UUID判断用户是否存在
func (u *userService) IsExistByUUID(uuid string) bool {
	var count int64
	u.db.Model(&dbmodel.UserDBModel{}).Where("uuid = ?", uuid).Count(&count)
	return count != 0
}

// IsExistByPhone 通过用户手机号判断用户是否存在
func (u *userService) IsExistByPhone(phone string) bool {
	var count int64
	u.db.Model(&dbmodel.UserDBModel{}).Where("phone = ?", phone).Count(&count)
	return count != 0
}

// IsExistByEmail 通过用户邮箱判断用户是否存在
func (u *userService) IsExistByEmail(email string) bool {
	var count int64
	u.db.Model(&dbmodel.UserDBModel{}).Where("email = ?", email).Count(&count)
	return count != 0
}

// GetUserByName 通过用户名获取用户信息
func (u *userService) GetUserByName(name string) dbmodel.UserDBModel {
	user := dbmodel.UserDBModel{}
	u.db.Where("username = ?", name).First(&user)
	return user
}

// GetUserByUUID 通过用户UUID获取用户信息
func (u *userService) GetUserByUUID(uid string) dbmodel.UserDBModel {
	user := dbmodel.UserDBModel{}
	u.db.Where("uuid = ?", uid).First(&user)
	return user
}

// GetUserByPhone 通过用户手机号获取用户信息
func (u *userService) GetUserByPhone(phone string) dbmodel.UserDBModel {
	user := dbmodel.UserDBModel{}
	u.db.Omit("password").Where("phone = ?", phone).First(&user)
	return user
}

// GetUserByEmail 通过用户邮箱获取用户信息
func (u *userService) GetUserByEmail(email string) dbmodel.UserDBModel {
	user := dbmodel.UserDBModel{}
	u.db.Omit("password").Where("email = ?", email).First(&user)
	return user
}
