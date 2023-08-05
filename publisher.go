package main

import (
	"fmt"
	"mqtt-publisher-subscriber-golang/consts"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(consts.Broker)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Error connecting to MQTT broker:", token.Error()))
	}

	for i := 1; i <= 10; i++ {
		message := fmt.Sprintf("Publishing message %d", i)
		token := client.Publish(consts.Topic, 0, false, message)
		token.Wait()

		fmt.Println("Published:", message)
		time.Sleep(1 * time.Second)
	}

	client.Disconnect(250)
}
