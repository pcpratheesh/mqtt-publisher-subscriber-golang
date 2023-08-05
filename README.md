# mqtt-publisher-subscriber-golang
This repository contains a simple example of an MQTT publisher and subscriber implemented in Golang using the popular "paho.mqtt.golang" library. The publisher publishes simulated temperature readings, and the subscriber subscribes to the topic to receive and display the messages.

**Key Features**:
- MQTT Publisher: Publishes messages to an MQTT broker at regular intervals.
- MQTT subscriber: Subscribes to the topic and displays received messages.
- Lightweight: Built with Golang and uses the lightweight MQTT protocol, making it efficient for IoT and M2M applications.


**Usage**:
1. Clone the repository: `git clone https://github.com/pcpratheesh/mqtt-publisher-subscriber-golang`
2. Make sure you have Golang installed on your system.
3. Install the required dependencies: `go get -u github.com/eclipse/paho.mqtt.golang`
4. Run the MQTT publisher: `go run publisher.go`
5. Run the MQTT subscriber: `go run subscriber.go`
