
# gonair (Go On Air)

gonair is a utility that detects when a macOS webcam is turned on or off and emits an MQTT message. This allows systems like Home Assistant to trigger automations based on the webcam's state.

## Purpose

With the rise of remote work and video conferencing, it's useful to have automations that respond to the state of your webcam. For instance, you might want to mute your smart speakers or change your room's lighting when you're on a call. gonair makes this possible by emitting MQTT messages when your webcam activates or deactivates.

## Getting Started

### Prerequisites

- Go
- An MQTT broker (e.g., Mosquitto, HiveMQ)
- macOS with webcam

### Installation

1. Clone the repository:
   ```
   git clone git@github.com:liftedkilt/gonair.git
   ```

2. Navigate to the project directory:
   ```
   cd gonair
   ```

3. Install the required Go packages:
   ```
   go get
   ```

### Configuration

Create a `.env` file in the root directory with your MQTT configurations:

```dotenv
LOG_POWER_STATE_ON='"VDCAssistant_Power_State" = On'
LOG_POWER_STATE=VDCAssistant_Power_State
LOG_PREDICATE='subsystem contains "com.apple.UVCExtension" and composedMessage contains "Post PowerLog"'
LOG_STYLE=syslog
MQTT_BROKER=tcp://192.168.0.100:1883
MQTT_CLIENT_ID=gonair-client
MQTT_PASSWORD=somesecretpassword
MQTT_TOPIC=home-assistant/gonair/state
MQTT_USERNAME=gonair
```

You can customize your log stream options based on your MacOS version and CPU architecture. I found the values used in the sample.env file [here](https://stackoverflow.com/questions/60535678/macos-detect-when-camera-is-turned-on-off)

### Running

Execute gonair with:
```
go run gonair.go
```

## License

gonair is licensed under the BSD 2-Clause License. Refer to the `LICENSE` file for details.
