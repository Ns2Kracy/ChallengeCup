package dbmodel

import "gorm.io/gorm"

type UserDBModel struct {
	gorm.Model
	ID        int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserName  string `gorm:"column:username;type:varchar(255);not null" json:"username"`
	Password  string `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Email     string `gorm:"column:email;type:varchar(255);not null" json:"email"`
	Phone     string `gorm:"column:phone;type:varchar(255);not null" json:"phone"`
	Role      int    `gorm:"column:role;type:int;not null" json:"role"`
	Avatar    string `gorm:"column:avatar;type:varchar(255);not null" json:"avatar"`
	CreatedAt int64  `gorm:"column:created_at;<-:create;autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;<-:create;<-:update;autoUpdateTime" json:"updated_at"`
}

func (UserDBModel) TableName() string {
	return "user"
}
