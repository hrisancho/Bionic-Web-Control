package mqtt

import (
	"Bionic-Web-Control/internal/config"
	main_logger "Bionic-Web-Control/internal/logger"
	"Bionic-Web-Control/proto/commands"
	"context"
	"crypto/tls"
	"crypto/x509"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	// TODO поправить на существующие протографы
)

var (
	//ControllerSubscriptionHandlerNotFound = errors.New("ControllerSubscriptionHandlerMQTT not found")
	notificationsTopicRegexp = regexp.MustCompile(`^controller/([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})/notifications/(.*)$`)
)

type ClientMQTT struct {
	logger *main_logger.Logger
	config config.Config
	client mqtt.Client
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

	return
}

func (clientMQTT *ClientMQTT) HandServoToAngle(uuid int, msg *commands.ServoGoToAngle) (err error) {
	// Передаем полностью обработанное сообщение
	msgMarshal, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	// На всякий случай сохраняем сообщения
	token := clientMQTT.client.Publish("robohand/"+strconv.Itoa(uuid)+"/commands/servo-go-to-angle", 0, true, msgMarshal)
	token.Wait()

	err = token.Error()
	if token.Error() != nil {
		err = token.Error()
		return
	}
	return nil
}

func (clientMQTT *ClientMQTT) HandServoLock(uuid int, msg *commands.ServoLock) (err error) {
	// Передаем полностью обработанное сообщение
	msgMarshal, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	// На всякий случай сохраняем сообщения
	token := clientMQTT.client.Publish("robohand/"+strconv.Itoa(uuid)+"/commands/servo-lock", 0, true, msgMarshal)
	token.Wait()

	err = token.Error()
	if token.Error() != nil {
		err = token.Error()
		return
	}
	return nil
}

func (clientMQTT *ClientMQTT) HandServoUnLock(uuid int, msg *commands.ServoUnLock) (err error) {
	// Передаем полностью обработанное сообщение
	msgMarshal, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	// На всякий случай сохраняем сообщения
	token := clientMQTT.client.Publish("robohand/"+strconv.Itoa(uuid)+"/commands/servo-unlock", 0, true, msgMarshal)
	token.Wait()

	err = token.Error()
	if token.Error() != nil {
		err = token.Error()
		return
	}
	return nil
}

func (clientMQTT *ClientMQTT) HandServoSmoothlyMove(uuid int, msg *commands.ServoSmoothlyMove) (err error) {
	// Передаем полностью обработанное сообщение
	msgMarshal, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	// На всякий случай сохраняем сообщения
	token := clientMQTT.client.Publish("robohand/"+strconv.Itoa(uuid)+"/commands/servo-smoothly-move", 0, true, msgMarshal)
	token.Wait()

	err = token.Error()
	if token.Error() != nil {
		err = token.Error()
		return
	}
	return nil
}
func (clientMQTT *ClientMQTT) MoveToTargetPressure(uuid int, msg *commands.MoveToTargetPressure) (err error) {
	// Передаем полностью обработанное сообщение
	msgMarshal, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	// На всякий случай сохраняем сообщения
	token := clientMQTT.client.Publish("robohand/"+strconv.Itoa(uuid)+"/commands/move-target-pressure", 0, true, msgMarshal)
	token.Wait()

	err = token.Error()
	if token.Error() != nil {
		err = token.Error()
		return
	}
	return nil
}

func (clientMQTT *ClientMQTT) ServoHoldGesture(uuid int, msg *commands.HoldGesture) (err error) {
	// Передаем полностью обработанное сообщение
	msgMarshal, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	// На всякий случай сохраняем сообщения
	token := clientMQTT.client.Publish("robohand/"+strconv.Itoa(uuid)+"/commands/hold-gesture", 0, true, msgMarshal)
	token.Wait()

	err = token.Error()
	if token.Error() != nil {
		err = token.Error()
		return
	}
	return nil
}

// TODO Удалить всё что не нужно находиться внизу
//
//func (clientMQTT *ClientMQTT) initControllerMqttMap(ctx context.Context) (err error) {
//	controllerList, err := clientMQTT.controllerStorage.All(ctx)
//	if err != nil {
//		return
//	}
//
//	for _, controller := range controllerList {
//		err = clientMQTT.PublishControllerUserAccessList(ctx, controller.UUID)
//		if err != nil {
//			clientMQTT.logger.Error("initControllerMqttMap PublishUserAccessList error", zap.String("uuid", controller.UUID), zap.Error(err))
//			continue
//		}
//	}
//
//	return
//}
//
//func (clientMQTT ClientMQTT) ControllerGetConfig(controllerUUID string) (controllerConfig controller_config.Config, err error) {
//	token := clientMQTT.client.Publish(fmt.Sprintf("controller/%s/config/get/request", controllerUUID), clientMQTT.config.MqttQOS, false, "")
//	token.Wait()
//	// TODO check return after clientMQTT.logger.Error("clientMQTT.
//	if err := token.Error(); err != nil {
//		clientMQTT.logger.Error("clientMQTT.ControllerGetConfig publishMQTT", zap.Error(err), zap.String("uuid", controllerUUID))
//	}
//
//	getConfigRequestCtx, getConfigRequestCtxCancel := context.WithTimeout(context.Background(), config.ControllerRequestTimeout)
//	getConfigRequestCtxTimeout := atomic.Bool{}
//	getConfigRequestCtxTimeout.Store(true)
//	defer getConfigRequestCtxCancel()
//
//	clientMQTT.client.Subscribe(fmt.Sprintf("controller/%s/config/get/response", controllerUUID), clientMQTT.config.MqttQOS, func(client mqtt.Client, msg mqtt.Message) {
//		err = proto.Unmarshal(msg.Payload(), &controllerConfig)
//		if err != nil {
//			return
//		}
//
//		getConfigRequestCtxTimeout.Store(false)
//		getConfigRequestCtxCancel()
//	})
//	<-getConfigRequestCtx.Done()
//	if err != nil {
//		return
//	}
//
//	if getConfigRequestCtxTimeout.Load() {
//		err = errors.New("Контроллер не отвечает")
//		return
//	}
//
//	err = token.Error()
//	return
//}
//
//func (clientMQTT ClientMQTT) ControllerSetConfig(controllerUUID string, controllerPassword string, controllerConfigStrJson string) (err error) {
//	hashSha256 := sha256.New()
//	hashSha256.Write([]byte(controllerPassword))
//	controllerPasswordHash := hashSha256.Sum(nil)
//
//	controllerConfig := controller_config.Config{
//		Config:                 controllerConfigStrJson,
//		ControllerPasswordHash: controllerPasswordHash,
//	}
//	log.Println(controllerPassword)
//	log.Printf("%x\n", controllerPasswordHash)
//	controllerConfigBytes, err := proto.Marshal(&controllerConfig)
//	if err != nil {
//		return
//	}
//	token := clientMQTT.client.Publish(fmt.Sprintf("controller/%s/config/set", controllerUUID), clientMQTT.config.MqttQOS, false, controllerConfigBytes)
//	token.Wait()
//	if err = token.Error(); err != nil {
//		clientMQTT.logger.Error("clientMQTT.ControllerSetControllerConfig publishMQTT", zap.Error(err), zap.String("uuid", controllerUUID))
//		return
//	}
//
//	requestCtx, requestCtxCancel := context.WithTimeout(context.Background(), time.Minute*5)
//	requestCtxTimeout := atomic.Bool{}
//	requestCtxTimeout.Store(true)
//	defer requestCtxCancel()
//
//	clientMQTT.client.Subscribe(fmt.Sprintf("controller/%s/notifications/config/+", controllerUUID), clientMQTT.config.MqttQOS, func(client mqtt.Client, msg mqtt.Message) {
//		defer func() {
//			requestCtxTimeout.Store(false)
//			requestCtxCancel()
//		}()
//
//		topicParts := strings.Split(msg.Topic(), "/")
//		if len(topicParts) == 0 {
//			err = errors.New("Неверный формат ответа контроллера")
//			return
//		}
//		responseStatus := topicParts[len(topicParts)-1]
//
//		switch responseStatus {
//		case "success":
//		case "error":
//			configSetErrorEvent := notifications.ConfigErrorEvent{}
//			err = proto.Unmarshal(msg.Payload(), &configSetErrorEvent)
//			if err != nil {
//				return
//			}
//
//			clientMQTT.logger.Error("clientMQTT.ControllerSetControllerConfig response error", zap.String("errorMsg", configSetErrorEvent.Error), zap.String("uuid", controllerUUID))
//			err = errors.New("Контроллер вернул ошибку при попытке изменения конфигурации")
//
//		default:
//			err = errors.New("Неверный статус ответа контроллера")
//		}
//	})
//	<-requestCtx.Done()
//	if err != nil {
//		return
//	}
//
//	if requestCtxTimeout.Load() {
//		err = errors.New("Контроллер не отвечает")
//		return
//	}
//
//	err = token.Error()
//	return
//}
//
//func (clientMQTT ClientMQTT) ControllerResetConfig(controllerUUID string, controllerPassword string) (err error) {
//	hashSha256 := sha256.New()
//	hashSha256.Write([]byte(controllerPassword))
//	controllerPasswordHash := hashSha256.Sum(nil)
//
//	request := controller_config.ResetConfig{
//		ControllerPasswordHash: controllerPasswordHash,
//	}
//
//	requestBytes, err := proto.Marshal(&request)
//	if err != nil {
//		return
//	}
//
//	token := clientMQTT.client.Publish(fmt.Sprintf("controller/%s/config/reset", controllerUUID), clientMQTT.config.MqttQOS, false, requestBytes)
//	token.Wait()
//	if err = token.Error(); err != nil {
//		clientMQTT.logger.Error("clientMQTT.ControllerSetControllerConfig publishMQTT", zap.Error(err), zap.String("uuid", controllerUUID))
//		return
//	}
//
//	requestCtx, requestCtxCancel := context.WithTimeout(context.Background(), time.Minute*5)
//	requestCtxTimeout := atomic.Bool{}
//	requestCtxTimeout.Store(true)
//	defer requestCtxCancel()
//
//	clientMQTT.client.Subscribe(fmt.Sprintf("controller/%s/notifications/config/+", controllerUUID), clientMQTT.config.MqttQOS, func(client mqtt.Client, msg mqtt.Message) {
//		defer func() {
//			requestCtxTimeout.Store(false)
//			requestCtxCancel()
//		}()
//
//		topicParts := strings.Split(msg.Topic(), "/")
//		if len(topicParts) == 0 {
//			err = errors.New("Неверный формат ответа контроллера")
//			return
//		}
//		responseStatus := topicParts[len(topicParts)-1]
//
//		switch responseStatus {
//		case "success":
//		case "error":
//			configSetErrorEvent := notifications.ConfigErrorEvent{}
//			err = proto.Unmarshal(msg.Payload(), &configSetErrorEvent)
//			if err != nil {
//				return
//			}
//
//			clientMQTT.logger.Error("clientMQTT.ControllerSetControllerConfig response error", zap.String("errorMsg", configSetErrorEvent.Error), zap.String("uuid", controllerUUID))
//			err = errors.New("Контроллер вернул ошибку при попытке изменения конфигурации")
//
//		default:
//			err = errors.New("Неверный статус ответа контроллера")
//		}
//	})
//	<-requestCtx.Done()
//	if err != nil {
//		return
//	}
//
//	if requestCtxTimeout.Load() {
//		err = errors.New("Контроллер не отвечает")
//		return
//	}
//
//	err = token.Error()
//	return
//}
//
//// ControllerOtaRequest Запрос обновления прошивки контроллера по воздуху
//func (clientMQTT ClientMQTT) ControllerOtaRequest(firmwareUuid string, controllerUUID string, controllerPassword string) (err error) {
//	hashSha256 := sha256.New()
//	hashSha256.Write([]byte(controllerPassword))
//	controllerPasswordHash := hashSha256.Sum(nil)
//
//	siteAddressWithoutProtocol := clientMQTT.config.SiteURL
//	index := strings.Index(siteAddressWithoutProtocol, "://")
//	if index != -1 {
//		siteAddressWithoutProtocol = siteAddressWithoutProtocol[index+3:]
//	}
//
//	request := controller_ota.Ota{
//		OtaUrl:                 fmt.Sprintf("%s://%s/api/v1/controller/%s/ota/", config.ControllerOtaFirmwareProtocol, siteAddressWithoutProtocol, firmwareUuid),
//		ControllerPasswordHash: controllerPasswordHash,
//	}
//
//	requestBytes, err := proto.Marshal(&request)
//	if err != nil {
//		return
//	}
//
//	token := clientMQTT.client.Publish(fmt.Sprintf("controller/%s/ota/request", controllerUUID), clientMQTT.config.MqttQOS, false, requestBytes)
//	token.Wait()
//	if err = token.Error(); err != nil {
//		clientMQTT.logger.Error("clientMQTT.ControllerSetControllerConfig publishMQTT", zap.Error(err), zap.String("uuid", controllerUUID))
//		return
//	}
//
//	requestCtx, requestCtxCancel := context.WithTimeout(context.Background(), time.Minute*5)
//	requestCtxTimeout := atomic.Bool{}
//	requestCtxTimeout.Store(true)
//	defer requestCtxCancel()
//
//	clientMQTT.client.Subscribe(fmt.Sprintf("controller/%s/notifications/ota/+", controllerUUID), clientMQTT.config.MqttQOS, func(client mqtt.Client, msg mqtt.Message) {
//		defer func() {
//			requestCtxTimeout.Store(false)
//			requestCtxCancel()
//		}()
//
//		topicParts := strings.Split(msg.Topic(), "/")
//		if len(topicParts) == 0 {
//			err = errors.New("Неверный формат ответа контроллера")
//			return
//		}
//		responseStatus := topicParts[len(topicParts)-1]
//
//		// TODO const static path from config
//		firmwarePath := filepath.Join("./static/", fmt.Sprintf("/controller-firmware/%s.bin", firmwareUuid))
//		switch responseStatus {
//		case "success":
//			if _, err = os.Stat(firmwarePath); err == nil {
//				err = os.Remove(firmwarePath)
//				if err != nil {
//					clientMQTT.logger.Error("Не удалось удалить файл прошивки",
//						zap.String("firmwarePath", firmwarePath),
//						zap.Error(err))
//				}
//			}
//		case "error":
//			err = errors.New("Контроллер вернул ошибку при обновления прошивки пользователя")
//
//		default:
//			err = errors.New("Неверный статус ответа контроллера")
//		}
//	})
//	<-requestCtx.Done()
//	if err != nil {
//		return
//	}
//
//	if requestCtxTimeout.Load() {
//		err = errors.New("Контроллер не отвечает")
//		return
//	}
//
//	err = token.Error()
//	return
//}
//
//func (clientMQTT ClientMQTT) DoorOpen(controllerUUID string) (err error) {
//	token := clientMQTT.client.Publish(fmt.Sprintf("controller/%s/control/door/open", controllerUUID), clientMQTT.config.MqttQOS, false, "")
//	token.Wait()
//	if err := token.Error(); err != nil {
//		clientMQTT.logger.Error("clientMQTT.DoorOpen publishMQTT", zap.Error(err), zap.String("uuid", controllerUUID))
//	}
//
//	// TODO await response
//
//	return token.Error()
//}
//
//func (clientMQTT ClientMQTT) DoorClose(controllerUUID string) (err error) {
//	token := clientMQTT.client.Publish(fmt.Sprintf("controller/%s/control/door/close", controllerUUID), clientMQTT.config.MqttQOS, false, "")
//	token.Wait()
//	if err := token.Error(); err != nil {
//		clientMQTT.logger.Error("clientMQTT.DoorClose publishMQTT", zap.Error(err), zap.String("uuid", controllerUUID))
//	}
//
//	return token.Error()
//}
//
//func (clientMQTT ClientMQTT) UserPhotoSelectListener() (err error) {
//	token := clientMQTT.client.Subscribe("backend/user_photo/select/request", clientMQTT.config.MqttQOS, func(client mqtt.Client, msg mqtt.Message) {
//		userPhotoMap, err := clientMQTT.userStorage.UserPhotoMap()
//
//		userPhotoSelectResponse := user_photo.SelectResponse{
//			UserPhotoList: make([]*user_photo.UserPhoto, 0, len(userPhotoMap)),
//			Error:         "",
//		}
//
//		for userId, userPhotoBytes := range userPhotoMap {
//			userPhotoSelectResponse.UserPhotoList = append(userPhotoSelectResponse.UserPhotoList, &user_photo.UserPhoto{
//				Id:    int32(userId),
//				Photo: userPhotoBytes,
//			})
//		}
//
//		userPhotoSelectResponseBytes, err := proto.Marshal(&userPhotoSelectResponse)
//		if err != nil {
//			clientMQTT.logger.Error("UserPhotoSelectListener: cant marshal response", zap.Error(err))
//			return
//		}
//
//		token := clientMQTT.client.Publish("backend/user_photo/select/response", clientMQTT.config.MqttQOS, false, userPhotoSelectResponseBytes)
//		token.Wait()
//		if err = token.Error(); err != nil {
//			clientMQTT.logger.Error("UserPhotoSelectListener: cant send response", zap.Error(err))
//			return
//		}
//	})
//
//	return token.Error()
//}
//
//func (clientMQTT *ClientMQTT) UserPhotoUpdate(userId uint32, photo []byte) (err error) {
//	clientMQTT.userPhotoUpdateMutex.Lock()
//	defer func() {
//		// TODO unsubscribe response check after requests
//		token := clientMQTT.client.Unsubscribe("backend/user_photo/update/response")
//		token.Wait()
//		if token.Error() != nil {
//			clientMQTT.logger.Error("UserPhotoUpdate: cant unsubscribe response topic")
//		}
//
//		clientMQTT.userPhotoUpdateMutex.Unlock()
//	}()
//
//	userPhotoUpdateRequestBytes, err := proto.Marshal(&user_photo.UpdateRequest{
//		UserPhoto: &user_photo.UserPhoto{
//			Id:    int32(userId),
//			Photo: photo,
//		},
//	})
//	if err != nil {
//		return
//	}
//
//	token := clientMQTT.client.Publish("backend/user_photo/update/request", clientMQTT.config.MqttQOS, false, userPhotoUpdateRequestBytes)
//	token.Wait()
//	if err = token.Error(); err != nil {
//		return
//	}
//
//	userPhotoUpdateRequestCtx, userPhotoUpdateRequestCtxCancel := context.WithTimeout(context.Background(), config.FaceRecognizerTimeout)
//	userPhotoUpdateRequestTimeout := atomic.Bool{}
//	userPhotoUpdateRequestTimeout.Store(true)
//	defer userPhotoUpdateRequestCtxCancel()
//
//	userPhotoUpdateResponse := user_photo.UpdateResponse{}
//	clientMQTT.client.Subscribe("backend/user_photo/update/response", clientMQTT.config.MqttQOS, func(client mqtt.Client, msg mqtt.Message) {
//		err = proto.Unmarshal(msg.Payload(), &userPhotoUpdateResponse)
//		if err != nil {
//			return
//		}
//
//		userPhotoUpdateRequestTimeout.Store(false)
//		userPhotoUpdateRequestCtxCancel()
//	})
//	<-userPhotoUpdateRequestCtx.Done()
//	if err != nil {
//		return
//	}
//
//	if userPhotoUpdateRequestTimeout.Load() {
//		return errors.New("Сервис распознавания лиц не отвечает")
//	}
//	if userPhotoUpdateResponse.GetError() != "" {
//		return errors.New(userPhotoUpdateResponse.GetError())
//	}
//
//	return
//}
