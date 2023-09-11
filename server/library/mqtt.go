package library

import (
	"fmt"

	"github.com/ChocolateAceCream/blog/global"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func InitMqttClient() mqtt.Client {
	config := global.CONFIG.Mqtt
	url := fmt.Sprintf("tcp://%s:%s", config.Host, config.Port)
	opts := mqtt.NewClientOptions().AddBroker(url)
	opts.SetClientID(config.ClientId)
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)
	client := mqtt.NewClient(opts)
	return client
}

func PublishMqttMsg(topic string, msg interface{}) {
	client := global.MQTT
	if !client.IsConnected() {
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	// Publish a message to a topic
	token := client.Publish(topic, 0, false, msg)
	token.Wait()
}

func SubscribeMqttMsg(topic string, cb mqtt.MessageHandler) {
	client := global.MQTT
	if !client.IsConnected() {
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	// Subscribe to a topic and print received messages
	token := client.Subscribe(topic, 0, cb)
	token.Wait()
}

func UnsubscribeMqttMsg(topic string) {
	client := global.MQTT
	if !client.IsConnected() {
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	// Unsubscribe from a topic
	token := client.Unsubscribe(topic)
	token.Wait()
}
