package service

import (
	"ChallengeCup/service/dbmodel"

	"gorm.io/gorm"
)

type MqttService interface {
	GetDataNow() dbmodel.MqttData
	GetTemperatureNow() float64
	GetHeartRateNow() uint32
	GetBloodOxygenNow() float64
	GetDataByTime(start string, end string) []dbmodel.MqttData
	GetTemperatureByTime(start string, end string) []float64
	GetHeartRateByTime(start string, end string) []uint32
	GetBloodOxygenByTime(start string, end string) []float64
}

// 自动检查未实现的接口

type mqttService struct {
	db *gorm.DB
}

func NewMqttService(db *gorm.DB) MqttService {
	return &mqttService{
		db: db,
	}
}

func (m *mqttService) GetDataNow() dbmodel.MqttData {
	var data dbmodel.MqttData
	m.db.Last(&data)
	return data
}

func (m *mqttService) GetTemperatureNow() float64 {
	var temperature float64
	m.db.Last(&temperature).Where("temperature != 0")
	return temperature
}

func (m *mqttService) GetHeartRateNow() uint32 {
	var heartRate uint32
	m.db.Last(&heartRate).Where("heart_rate != 0")
	return heartRate
}

func (m *mqttService) GetBloodOxygenNow() float64 {
	var bloodOxygen float64
	m.db.Last(&bloodOxygen).Where("blood_oxygen != 0")
	return bloodOxygen
}

func (m *mqttService) GetDataByTime(start string, end string) []dbmodel.MqttData {
	data := []dbmodel.MqttData{}
	m.db.Where("created_at >= ? AND created_at <= ?", start, end).Find(&data)
	return data
}

func (m *mqttService) GetTemperatureByTime(start string, end string) []float64 {
	var data []float64
	m.db.Where("created_at >= ? AND created_at <= ?", start, end).Find(&data).Where("temperature != 0")
	return data
}

func (m *mqttService) GetHeartRateByTime(start string, end string) []uint32 {
	var data []uint32
	m.db.Where("created_at >= ? AND created_at <= ?", start, end).Find(&data).Where("heart_rate != 0")
	return nil
}

func (m *mqttService) GetBloodOxygenByTime(start string, end string) []float64 {
	var data []float64
	m.db.Where("created_at >= ? AND created_at <= ?", start, end).Find(&data).Where("blood_oxygen != 0")
	return data
}
