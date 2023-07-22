package main

import (
	"fmt"
	"message-broker/helper"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	var brokerURI = "tcp://localhost:1883"
	options := mqtt.NewClientOptions()
	options.AddBroker(brokerURI)
	options.SetClientID("go_mqtt_sender")
	options.SetDefaultPublishHandler(helper.MessagePubHandler)
	options.OnConnect = helper.ConnectHandler
	options.OnConnectionLost = helper.ConnectionLostHandler

	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	topic := "topic/secret"
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("counting %d", i)
		token = client.Publish(topic, 0, false, text)
		token.Wait()
		time.Sleep(1 * time.Second)
	}

	client.Disconnect(100)
}
