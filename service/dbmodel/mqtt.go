package dbmodel

type MqttData struct {
	ID          uint32  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Temperature float64 `gorm:"column:temperature;type:float;not null" json:"temperature"`
	HeartRate   uint32  `gorm:"column:heart_rate;type:float;not null" json:"heart_rate"`
	BloodOxygen float64 `gorm:"column:blood_oxygen;type:float;not null" json:"blood_oxygen"`
	CreatedAt   int64   `gorm:"column:created_at;<-:create;autoCreateTime" json:"created_at"`
}

func (MqttData) TableName() string {
	return "mqtt_data"
}
