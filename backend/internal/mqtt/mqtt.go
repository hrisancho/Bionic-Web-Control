package mqtt

import (
	"Bionic-Web-Control/internal/config"
	main_logger "Bionic-Web-Control/internal/logger"
	"Bionic-Web-Control/proto/imu"
	"Bionic-Web-Control/proto/potentiometer"
	"Bionic-Web-Control/proto/servo"
	"Bionic-Web-Control/proto/straingauge"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	notificationsTopicRegexp = regexp.MustCompile(`^controller/([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})/notifications/(.*)$`)
)

type ClientMQTT struct {
	logger *main_logger.Logger
	config config.Config
	client mqtt.Client
	// uuid в качестве ключа
	StorageRawImu     map[string]*imu.IMU
	StorageProcessImu map[string]*imu.ResultIMU
	// uuid в качестве первого ключа, а качестве второго палец
	StorageStrainGauge map[string]map[string]*staingauge.StrainGuage
	// uuid в качестве первого ключа, а в качестве второго позиция сервопривод
	StorageServoInfo map[string]map[string]*servo.Servo
	// uuid в качестве первого ключа, а в качестве второго палец, а в качестве третьего позиция потенциометра
	StoragePotentiometerAngle map[string]map[string]map[string]*potentiometer.Potentiometer
}

func NewClientMQTT(
	ctx context.Context,
	logger *main_logger.Logger,
	config config.Config) (clientMQTT *ClientMQTT, err error) {

	clientMQTT = &ClientMQTT{
		logger: logger,
		config: config,
	}

	opts := mqtt.NewClientOptions().
		AddBroker(config.MqttBrokerAddr).
		// id клиента будет зависеть от времени обращения сервера
		SetClientID("backend-" + time.Now().String()).
		SetKeepAlive(config.MqttKeepAlive).
		SetPingTimeout(config.MqttPingTimeout).
		SetAutoReconnect(true).
		SetConnectionLostHandler(func(c mqtt.Client, err error) {
			clientMQTT.logger.Warn("MQTT connection lost", zap.Error(err))
		}).
		SetReconnectingHandler(func(c mqtt.Client, options *mqtt.ClientOptions) {
			clientMQTT.logger.Warn("MQTT reconnecting", zap.Error(err))
		})

	opts.OnConnect = func(client mqtt.Client) {
		clientMQTT.logger.Info("MQTT client connected")
	}
	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		clientMQTT.logger.Info("MQTT Connect lost : ", zap.Error(err))
	}
	// Проверяет, есть ли с начала есть префикс означающий защищённое соединение
	if strings.HasPrefix(config.MqttBrokerAddr, "ssl") {
		cert, err := tls.LoadX509KeyPair("./config/mqtt-tls/client.crt", "./config/mqtt-tls/client.key")
		if err != nil {
			clientMQTT.logger.Fatal("MQTT: SSL enabled, but can't read client.crt/client.key/cert.pem (in ./config/mqtt-tls folder)", zap.Error(err))
		}

		certpool := x509.NewCertPool()
		ca, err := os.ReadFile("./config/mqtt-tls/cert.pem")
		if err != nil {
			clientMQTT.logger.Fatal("MQTT: SSL enabled, but can't read cert.pem", zap.Error(err))
		}

		certpool.AppendCertsFromPEM(ca)
		opts.SetTLSConfig(&tls.Config{
			RootCAs:      certpool,
			ClientAuth:   tls.RequestClientCert,
			Certificates: []tls.Certificate{cert},
		})
		clientMQTT.logger.Info("MQTT: SSL enabled")
	}

	if len(config.MqttUsername) > 0 {
		opts.SetUsername(config.MqttUsername)
	}

	if len(config.MqttPassword) > 0 {
		opts.SetPassword(config.MqttPassword)
	}

	clientMQTT.client = mqtt.NewClient(opts)
	if token := clientMQTT.client.Connect(); token.Wait() && token.Error() != nil {
		err = token.Error()
		return
	}
	// Инициализация хранилищ, для быстрого сохранения и получения информации из них
	clientMQTT.StorageRawImu = make(map[string]*imu.IMU)
	clientMQTT.StorageProcessImu = make(map[string]*imu.ResultIMU)
	clientMQTT.StorageStrainGauge = make(map[string]map[string]*staingauge.StrainGuage)
	clientMQTT.StorageServoInfo = make(map[string]map[string]*servo.Servo)
	clientMQTT.StoragePotentiometerAngle = make(map[string]map[string]map[string]*potentiometer.Potentiometer)
	return
}

// TODO стоит ли сравнивать предыдущее значение с нынешенем только что полученным для дальнейшего изменения (WebSocket)
func (clientMQTT *ClientMQTT) InitSubscribeToUUID(uuid string) (err error) {
	clientMQTT.client.Subscribe("robohand/"+uuid+"/monitoring/IMU/raw-data", clientMQTT.config.MqttQOS, clientMQTT.ImuRawData)
	clientMQTT.client.Subscribe("robohand/"+uuid+"/monitoring/IMU/processed-data", clientMQTT.config.MqttQOS, clientMQTT.ImuProcessData)
	clientMQTT.client.Subscribe("robohand/"+uuid+"/monitoring/strain_gauge/pressure-at-fingertips", clientMQTT.config.MqttQOS, clientMQTT.StrainGaugeFingertips)
	clientMQTT.client.Subscribe("robohand/"+uuid+"/monitoring/servo/info", clientMQTT.config.MqttQOS, clientMQTT.ServoInfo)
	clientMQTT.client.Subscribe("robohand/"+uuid+"/monitoring/potentiometer/angle-measurement", clientMQTT.config.MqttQOS, clientMQTT.PotentiomAngle)
	return
}
func (clientMQTT *ClientMQTT) ImuRawData(client mqtt.Client, msg mqtt.Message) {
	uuid := TopicToUUID(msg.Topic())
	reqest := msg.Payload()[:]
	imuMsg := &imu.IMU{}
	err := proto.Unmarshal(reqest, imuMsg)
	if err != nil {
		clientMQTT.logger.Warn("MQTT: ImuRawData", zap.Error(err))
	}
	fmt.Println(imuMsg)
	clientMQTT.StorageRawImu[uuid] = imuMsg
	fmt.Println(clientMQTT.StorageRawImu)
}

func (clientMQTT *ClientMQTT) ImuProcessData(client mqtt.Client, msg mqtt.Message) {
	uuid := TopicToUUID(msg.Topic())
	reqest := msg.Payload()[:]
	imuProcMsg := &imu.ResultIMU{}
	err := proto.Unmarshal(reqest, imuProcMsg)
	if err != nil {
		clientMQTT.logger.Warn("MQTT: ImuProcessData", zap.Error(err))
	}
	fmt.Println(imuProcMsg)
	clientMQTT.StorageProcessImu[uuid] = imuProcMsg
	fmt.Println(clientMQTT.StorageProcessImu)
}

func (clientMQTT *ClientMQTT) StrainGaugeFingertips(client mqtt.Client, msg mqtt.Message) {
	uuid := TopicToUUID(msg.Topic())
	reqest := msg.Payload()[:]
	stnFingertips := &staingauge.StrainGuage{}
	err := proto.Unmarshal(reqest, stnFingertips)
	if err != nil {
		clientMQTT.logger.Warn("MQTT: StrainGaugeFingertips", zap.Error(err))
	}

	fmt.Println(stnFingertips)

	if clientMQTT.StorageStrainGauge[uuid] == nil {
		clientMQTT.StorageStrainGauge[uuid] = make(map[string]*staingauge.StrainGuage)
	}
	clientMQTT.StorageStrainGauge[uuid][stnFingertips.Finger.String()] = stnFingertips
	fmt.Println(clientMQTT.StorageStrainGauge)
}

func (clientMQTT *ClientMQTT) ServoInfo(client mqtt.Client, msg mqtt.Message) {
	uuid := TopicToUUID(msg.Topic())
	reqest := msg.Payload()[:]
	srvInfo := &servo.Servo{}
	err := proto.Unmarshal(reqest, srvInfo)
	if err != nil {
		clientMQTT.logger.Warn("MQTT: ServoInfo", zap.Error(err))
	}
	fmt.Println(srvInfo)

	if clientMQTT.StorageServoInfo[uuid] == nil {
		clientMQTT.StorageServoInfo[uuid] = make(map[string]*servo.Servo)
	}
	clientMQTT.StorageServoInfo[uuid][srvInfo.Servo.String()] = srvInfo
	fmt.Println(clientMQTT.StorageServoInfo)
}

func (clientMQTT *ClientMQTT) PotentiomAngle(client mqtt.Client, msg mqtt.Message) {
	uuid := TopicToUUID(msg.Topic())
	reqest := msg.Payload()[:]
	potentiometAngle := &potentiometer.Potentiometer{}
	err := proto.Unmarshal(reqest, potentiometAngle)
	if err != nil {
		clientMQTT.logger.Warn("MQTT: PotentiomAngle", zap.Error(err))
	}
	fmt.Println(potentiometAngle)
	if clientMQTT.StoragePotentiometerAngle[uuid] == nil {
		clientMQTT.StoragePotentiometerAngle[uuid] = make(map[string]map[string]*potentiometer.Potentiometer)
	}
	if clientMQTT.StoragePotentiometerAngle[uuid][potentiometAngle.Finger.String()] == nil {
		clientMQTT.StoragePotentiometerAngle[uuid][potentiometAngle.Finger.String()] = make(map[string]*potentiometer.Potentiometer)
	}
	clientMQTT.StoragePotentiometerAngle[uuid][potentiometAngle.Finger.String()][potentiometAngle.Positoin.String()] = potentiometAngle
	fmt.Println(clientMQTT.StoragePotentiometerAngle)
}

// Нахождение uuid протеза через топик сообщения
func TopicToUUID(topic string) string {
	return strings.Split(topic, "/")[1]

}
