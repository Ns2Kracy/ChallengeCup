package dbmodel

import "gorm.io/gorm"

type UserDBModel struct {
	gorm.Model
	ID             int    `gorm:"column:id;primaryKey;autoIncrement;->:false;<-:create" json:"id"`
	UUID           string `gorm:"column:uuid;type:varchar(30);not null" json:"uuid"`
	UserName       string `gorm:"column:username;type:varchar(255);not null" json:"username"`
	Password       string `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Email          string `gorm:"column:email;type:varchar(255);not null" json:"email"`
	Phone          string `gorm:"column:phone;type:varchar(30);not null" json:"phone"`
	IsEmailActived bool   `gorm:"column:is_email_actived;type:boolean;default:false;not null" json:"is_email_actived"`
	IsPhoneActived bool   `gorm:"column:is_phone_actived;type:boolean;default:false;not null" json:"is_phone_actived"`
	Avatar         string `gorm:"column:avatar;type:varchar(255);not null" json:"avatar"`
	EmailActivedAt int64  `gorm:"column:email_actived_at;type:bigint;default:0;not null" json:"email_actived_at"`
	PhoneActivedAt int64  `gorm:"column:phone_actived_at;type:bigint;default:0;not null" json:"phone_actived_at"`
	CreatedAt      int64  `gorm:"column:created_at;<-:create;autoCreateTime" json:"created_at"`
	UpdatedAt      int64  `gorm:"column:updated_at;<-:create;<-:update;autoUpdateTime" json:"updated_at"`
}

func (UserDBModel) TableName() string {
	return "user"
}
