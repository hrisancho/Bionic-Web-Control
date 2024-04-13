package mqtt

import (
	"Bionic-Web-Control/proto/commands"
	"google.golang.org/protobuf/proto"
)

// Ниже идут методы для взаимодействие с командами для контроллера
func (clientMQTT *ClientMQTT) HandServoToAngle(uuid string, msg *commands.ServoGoToAngle) (err error) {
	// Передаем полностью обработанное сообщение
	msgMarshal, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	// На всякий случай сохраняем сообщения
	token := clientMQTT.client.Publish("robohand/"+uuid+"/commands/servo-go-to-angle", clientMQTT.config.MqttQOS, true, msgMarshal)
	token.Wait()

	err = token.Error()
	if token.Error() != nil {
		err = token.Error()
		return
	}
	return nil
}

func (clientMQTT *ClientMQTT) HandServoLock(uuid string, msg *commands.ServoLock) (err error) {
	// Передаем полностью обработанное сообщение
	msgMarshal, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	// На всякий случай сохраняем сообщения
	token := clientMQTT.client.Publish("robohand/"+uuid+"/commands/servo-lock", clientMQTT.config.MqttQOS, true, msgMarshal)
	token.Wait()

	err = token.Error()
	if token.Error() != nil {
		err = token.Error()
		return
	}
	return nil
}

func (clientMQTT *ClientMQTT) HandServoUnLock(uuid string, msg *commands.ServoUnLock) (err error) {
	// Передаем полностью обработанное сообщение
	msgMarshal, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	// На всякий случай сохраняем сообщения
	token := clientMQTT.client.Publish("robohand/"+uuid+"/commands/servo-unlock", clientMQTT.config.MqttQOS, true, msgMarshal)
	token.Wait()

	err = token.Error()
	if token.Error() != nil {
		err = token.Error()
		return
	}
	return nil
}

func (clientMQTT *ClientMQTT) HandServoSmoothlyMove(uuid string, msg *commands.ServoSmoothlyMove) (err error) {
	// Передаем полностью обработанное сообщение
	msgMarshal, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	// На всякий случай сохраняем сообщения
	token := clientMQTT.client.Publish("robohand/"+uuid+"/commands/servo-smoothly-move", clientMQTT.config.MqttQOS, true, msgMarshal)
	token.Wait()

	err = token.Error()
	if token.Error() != nil {
		err = token.Error()
		return
	}
	return nil
}
func (clientMQTT *ClientMQTT) MoveToTargetPressure(uuid string, msg *commands.MoveToTargetPressure) (err error) {
	// Передаем полностью обработанное сообщение
	msgMarshal, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	// На всякий случай сохраняем сообщения
	token := clientMQTT.client.Publish("robohand/"+uuid+"/commands/move-target-pressure", clientMQTT.config.MqttQOS, true, msgMarshal)
	token.Wait()

	err = token.Error()
	if token.Error() != nil {
		err = token.Error()
		return
	}
	return nil
}

func (clientMQTT *ClientMQTT) ServoHoldGesture(uuid string, msg *commands.HoldGesture) (err error) {
	// Передаем полностью обработанное сообщение
	msgMarshal, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	// На всякий случай сохраняем сообщения
	token := clientMQTT.client.Publish("robohand/"+uuid+"/commands/hold-gesture", clientMQTT.config.MqttQOS, true, msgMarshal)
	token.Wait()

	err = token.Error()
	if token.Error() != nil {
		err = token.Error()
		return
	}
	return nil
}
