# mqtt-producer-consumer-golang
This repository contains a simple example of an MQTT producer and consumer implemented in Golang using the popular "paho.mqtt.golang" library. The producer publishes simulated temperature readings, and the consumer subscribes to the topic to receive and display the messages.

**Key Features**:
- MQTT Producer: Publishes messages to an MQTT broker at regular intervals.
- MQTT Consumer: Subscribes to the topic and displays received messages.
- Lightweight: Built with Golang and uses the lightweight MQTT protocol, making it efficient for IoT and M2M applications.


**Usage**:
1. Clone the repository: `git clone https://github.com/pcpratheesh/mqtt-producer-consumer-golang`
2. Make sure you have Golang installed on your system.
3. Install the required dependencies: `go get -u github.com/eclipse/paho.mqtt.golang`
4. Run the MQTT producer: `go run producer.go`
5. Run the MQTT consumer: `go run consumer.go`
