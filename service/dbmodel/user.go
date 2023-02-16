package dbmodel

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type UserDBModel struct {
	gorm.Model
	ID             int    `gorm:"column:id;primaryKey;autoIncrement;" json:"-"`
	UUID           uint32 `gorm:"column:uuid;type:bigint;not null" json:"uuid"`
	UserName       string `gorm:"column:username;type:varchar(255);not null" json:"username"`
	Password       string `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Phone          string `gorm:"column:phone;type:varchar(30)" json:"phone"`
	IsPhoneActived bool   `gorm:"column:is_phone_actived;type:boolean;default:false" json:"is_phone_actived"`
	Avatar         string `gorm:"column:avatar;type:varchar(255);" json:"avatar"`
	PhoneActivedAt int64  `gorm:"column:phone_actived_at;<-:create;<-:update;autoUpdateTime" json:"phone_actived_at"`
	CreatedAt      int64  `gorm:"column:created_at;<-:create;autoCreateTime" json:"created_at"`
	UpdatedAt      int64  `gorm:"column:updated_at;<-:create;<-:update;autoUpdateTime" json:"updated_at"`
}

func (UserDBModel) TableName() string {
	return "user"
}

type EmergencyPhone []string

func (ep *EmergencyPhone) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, ep)
}

func (ep EmergencyPhone) Value() (driver.Value, error) {
	return json.Marshal(ep)
}
