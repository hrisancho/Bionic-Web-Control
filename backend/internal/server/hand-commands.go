package server

import (
	"Bionic-Web-Control/proto/commands"
	"Bionic-Web-Control/proto/shared"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (server *Server) handInitSubscribeToUUID(c *fiber.Ctx) error {
	// TODO доделать, как минимум написать reqest
	uuid := c.Params("uuid")
	err := server.clientMQTT.InitSubscribeToUUID(uuid)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return nil
}

// TODO добавить тупой request без проверки правильность

type ServoGoToAngle struct {
	Angle float32 `json:"angle"`
}

// uuid и srvo_position указан
//
//	{
//		"angle" : 56.5, это угол передаётся через json
//	}
func (server *Server) handServoToAngle(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	servoId := c.Params("servo_id")
	srv, fun := shared.ServoPosition_value[servoId]
	if !fun {
		return fiber.NewError(fiber.StatusBadRequest, "Сервопривод указан не правильно")
	}
	srvPosition := shared.ServoPosition(uint32(srv))

	jsonRequest := new(ServoGoToAngle)
	err := c.BodyParser(jsonRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	msg := &commands.ServoGoToAngle{
		Servo: srvPosition,
		Angle: jsonRequest.Angle,
	}

	fmt.Println(msg)

	err = server.clientMQTT.HandServoToAngle(uuid, msg)
	if err != nil {
		return err
	}
	return nil
}

func (server *Server) handServoLock(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	servoId := c.Params("servo_id")
	srv, fun := shared.ServoPosition_value[servoId]
	if !fun {
		return fiber.NewError(fiber.StatusBadRequest, "Сервопривод указан не правильно")
	}
	srvPosition := shared.ServoPosition(uint32(srv))

	msg := &commands.ServoLock{
		Servo: srvPosition,
	}

	fmt.Println(msg)

	err := server.clientMQTT.HandServoLock(uuid, msg)
	if err != nil {
		return err
	}
	return nil
}

func (server *Server) handServoUnlock(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	servoId := c.Params("servo_id")
	srv, fun := shared.ServoPosition_value[servoId]
	if !fun {
		return fiber.NewError(fiber.StatusBadRequest, "Сервопривод указан не правильно")
	}
	srvPosition := shared.ServoPosition(uint32(srv))

	msg := &commands.ServoUnLock{
		Servo: srvPosition,
	}

	fmt.Println(msg)

	err := server.clientMQTT.HandServoUnLock(uuid, msg)
	if err != nil {
		return err
	}
	return nil
}

// //Постепенное движение сервоприводом

func (server *Server) handServoSmoothlyMove(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	servoId := c.Params("servo_id")
	srv, fun := shared.ServoPosition_value[servoId]
	if !fun {
		return fiber.NewError(fiber.StatusBadRequest, "Сервопривод указан не правильно")
	}
	srvPosition := shared.ServoPosition(uint32(srv))

	jsonRequest := new(ServoSmoothlyMove)
	err := c.BodyParser(jsonRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	easing, fun := commands.Easings_value[jsonRequest.Easing]
	if !fun {
		return fiber.NewError(fiber.StatusBadRequest, "Функция, задающая движения скорости от времени указана не правильно")
	}

	msg := &commands.ServoSmoothlyMove{
		Servo:       srvPosition,
		Easing:      commands.Easings(uint32(easing)),
		Speed:       jsonRequest.Speed,
		TargetAngle: jsonRequest.TargetAngle,
	}

	fmt.Println(msg)

	err = server.clientMQTT.HandServoSmoothlyMove(uuid, msg)
	if err != nil {
		return err
	}
	return nil
}

type ServoSmoothlyMove struct {
	Easing      string  `json:"easing"`
	Speed       float32 `json:"speed"`
	TargetAngle float32 `json:"targetAngle"`
}

func (server *Server) handMoveToTargetPressure(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	servoId := c.Params("servo_id")
	srv, fun := shared.ServoPosition_value[servoId]
	if !fun {
		return fiber.NewError(fiber.StatusBadRequest, "Сервопривод указан не правильно")
	}
	srvPosition := shared.ServoPosition(uint32(srv))

	fingerId := c.Params("finger_id")
	fng, fun := shared.Finger_value[fingerId]
	if !fun {
		return fiber.NewError(fiber.StatusBadRequest, "Палец указан не правильно")
	}
	fngPosition := shared.Finger(uint32(fng))

	msg := &commands.MoveToTargetPressure{
		Servo:  srvPosition,
		Finger: fngPosition,
	}

	fmt.Println(msg)

	err := server.clientMQTT.MoveToTargetPressure(uuid, msg)
	if err != nil {
		return err
	}
	return nil
}

type HoldGesture struct {
	Gesture   string `json:"gesture"`
	Duration  uint32 `json:"duration"`
	Permanent bool   `json:"permanent"`
}

func (server *Server) handServoHoldGesture(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	jsonRequest := new(HoldGesture)
	err := c.BodyParser(jsonRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	gts, fun := commands.Gestures_value[jsonRequest.Gesture]
	if !fun {
		return fiber.NewError(fiber.StatusBadRequest, "Жест указан не правильно")
	}
	gtsPosition := commands.Gestures(gts)

	msg := &commands.HoldGesture{
		Gesture:   gtsPosition,
		Duration:  jsonRequest.Duration,
		Permanent: jsonRequest.Permanent,
	}

	fmt.Println(msg)

	err = server.clientMQTT.ServoHoldGesture(uuid, msg)
	if err != nil {
		return err
	}
	return nil
}
