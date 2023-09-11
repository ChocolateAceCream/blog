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
	opts.SetClientID("blog")
	opts.SetUsername("superadmin")
	opts.SetPassword("123qwe!@#QWE")
	client := mqtt.NewClient(opts)
	return client
}

func PublishMqttMsg(topic string, msg string) {
	client := global.MQTT
	fmt.Println("------stat published to topic----", client.IsConnected())
	if !client.IsConnected() {
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			fmt.Println("------token.Error()-----", token.Error())
			panic(token.Error())
		} else {
			fmt.Println("------b4 published to topic----")

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
	fmt.Println("--UnsubscribeMqttMsg----")
	token.Wait()
}
