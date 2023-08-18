package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

type MQTTConfig struct {
	client mqtt.Client
	topic  string
}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(os.Getenv("MQTT_BROKER"))
	opts.SetClientID(os.Getenv("MQTT_CLIENT_ID"))
	opts.SetUsername(os.Getenv("MQTT_USERNAME"))
	opts.SetPassword(os.Getenv("MQTT_PASSWORD"))

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	defer client.Disconnect(250)

	config := MQTTConfig{
		client: client,
		topic:  os.Getenv("MQTT_TOPIC"),
	}

	config.monitorLogs()
}

func (m MQTTConfig) publishMessage(message string) {
	token := m.client.Publish(m.topic, 0, false, message)
	token.Wait()
}

func (m MQTTConfig) monitorLogs() {
	cmd := exec.Command("log", "stream", "--style", "syslog", "--predicate", `subsystem contains "com.apple.UVCExtension" and composedMessage contains "Post PowerLog"`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		m.handleLogEvent(line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func (m MQTTConfig) handleLogEvent(log string) {
	if strings.Contains(log, "VDCAssistant_Power_State") {
		// Extract the power state
		powerState := extractPowerState(log)
		m.processPowerState(powerState)
	}
}

func extractPowerState(log string) bool {
	return strings.Contains(log, `VDCAssistant_Power_State" = On`)
}

func (m MQTTConfig) processPowerState(state bool) {
	if state {
		fmt.Println("Camera turned on")
		m.publishMessage("ON")
	} else {
		fmt.Println("Camera turned off")
		m.publishMessage("OFF")
	}
}
