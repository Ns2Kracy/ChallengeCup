package service

import (
	"ChallengeCup/service/dbmodel"

	"gorm.io/gorm"
)

type UserService interface {
	// Warn: 注销用户
	DeleteByUUID(uuid string)

	// New: 注册新用户
	NewUser(user dbmodel.UserDBModel) dbmodel.UserDBModel

	// Updates: 更新用户信息
	UpdateUser(user dbmodel.UserDBModel)
	UpdatePassword(uuid string, password string)
	UploadAvatar(uuid string, avatar string)

	// Checks
	IsExistByUUID(uuid string) bool
	CheckPhone(phone string) bool

	// Get: 获取用户信息
	GetUserByName(name string) dbmodel.UserDBModel
	GetUserByUUID(uuid string) dbmodel.UserDBModel
	GetUserByPhone(phone string) dbmodel.UserDBModel
}

var _ UserService = (*userService)(nil)

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

// IsExistByUUID 通过用户UUID判断用户是否存在
func (u *userService) IsExistByUUID(uuid string) bool {
	var count int64
	u.db.Model(&dbmodel.UserDBModel{}).Where("uuid = ?", uuid).Count(&count)
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
	u.db.Omit("password").Where("uuid = ?", uid).First(&user)
	return user
}

// GetUserByPhone 通过用户手机号获取用户信息
func (u *userService) GetUserByPhone(phone string) dbmodel.UserDBModel {
	user := dbmodel.UserDBModel{}
	u.db.Where("phone = ?", phone).First(&user)
	return user
}

// CheckPhone 检查手机号是否可用
func (u *userService) CheckPhone(phone string) bool {
	var count int64
	u.db.Model(&dbmodel.UserDBModel{}).Where("phone = ?", phone).Count(&count)
	return count == 0
}
