package main

import (
	"fmt"
	"mqtt-producer-consumer-golang/consts"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(consts.Broker)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Error connecting to MQTT broker:", token.Error())
		os.Exit(1)
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
