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
	options.SetClientID("go_mqtt_receiver")
	options.SetDefaultPublishHandler(helper.MessagePubHandler)
	options.OnConnect = helper.ConnectHandler
	options.OnConnectionLost = helper.ConnectionLostHandler

	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	topic := "topic/secret"
	token = client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic %s\n", topic)
	time.Sleep(30 * time.Second)
	client.Disconnect(100)
}
