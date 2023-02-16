package mqtt

import (
	"encoding/json"

	"ChallengeCup/config"
	"ChallengeCup/dao"
	"ChallengeCup/service/dbmodel"
	log "ChallengeCup/utils/logger"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	Client   mqtt.Client
	callback = func(client mqtt.Client, msg mqtt.Message) {
		var message dbmodel.MqttData
		err := json.Unmarshal(msg.Payload(), &message)
		if err != nil {
			log.Errorf("MQTT 反序列化失败", err)
			return
		}
		dao.DB.Create(&message)
	}
)

func RunMqttClient() {
	conf := config.LoadConfig().Mqtt
	options := mqtt.NewClientOptions()
	options.AddBroker(conf.Broker).
		SetClientID(conf.ClientId).
		SetUsername(conf.Username).
		SetPassword(conf.Password).
		SetWill("server_will", "lose_connect", 2, false)
	Client = mqtt.NewClient(options)
	if token := Client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT Connect Error: ", token.Error())
	}
	log.Infof("MQTT Connect Success")
}

func Subscribe(topic string, callback mqtt.MessageHandler) {
	if token := Client.Subscribe(topic, 0, callback); token.Wait() && token.Error() != nil {
		log.Errorf("MQTT 订阅%s失败", topic, token.Error())
	}
	log.Infof("MQTT 订阅%s成功", topic)
}

func Unsubscribe(topic string) {
	if token := Client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		log.Errorf("MQTT 取消订阅%s失败", topic, token.Error())
	}
	log.Infof("MQTT 取消订阅%s成功", topic)
}

func Publish(topic string, payload interface{}) {
	if token := Client.Publish(topic, 0, true, payload); token.Wait() && token.Error() != nil {
		log.Errorf("MQTT 向%s发布失败", topic, token.Error())
	}
	log.Infof("MQTT 向%s发布成功", topic)
}

func Reply(topic string, replay interface{}) {
	marshal, err := json.Marshal(replay)
	if err != nil {
		log.Errorf("MQTT 序列化失败", err)
		return
	}
	Publish(topic, string(marshal))
}

func InitMqtt() {
	RunMqttClient()
	Subscribe("test/#", callback)
}
