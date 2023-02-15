package mqtt

import (
	"encoding/json"

	"ChallengeCup/config"
	"ChallengeCup/dao"
	log "ChallengeCup/utils/logger"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	Client   mqtt.Client
	callback = func(client mqtt.Client, msg mqtt.Message) {
		// * TODO: 处理消息
		dao.DB.Create("")
		log.Info("MQTT 收到消息: %s", msg.Payload())
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
	log.Info("MQTT Connect Success")
}

func Subscribe(topic string, callback mqtt.MessageHandler) {
	if token := Client.Subscribe(topic, 0, callback); token.Wait() && token.Error() != nil {
		log.Error("MQTT 订阅%s失败", topic, token.Error())
	}
}

func Unsubscribe(topic string) {
	if token := Client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		log.Error("MQTT 取消订阅%s失败", topic, token.Error())
	}
}

func Publish(topic string, payload interface{}) {
	if token := Client.Publish(topic, 0, true, payload); token.Wait() && token.Error() != nil {
		log.Error("MQTT 向%s发布失败", topic, token.Error())
	}
}

func Reply(topic string, replay interface{}) {
	marshal, err := json.Marshal(replay)
	if err != nil {
		log.Error("MQTT 序列化失败", err)
		return
	}
	Publish(topic, string(marshal))
}
