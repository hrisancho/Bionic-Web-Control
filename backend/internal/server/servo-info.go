package server

import (
	"Bionic-Web-Control/proto/shared"
	"github.com/gofiber/fiber/v2"
)

func (server *Server) handServoInfoByServoId(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	servoId := c.Params("servo_id")

	srv, fun := shared.ServoPosition_value[servoId]
	if !fun {
		return fiber.NewError(fiber.StatusBadRequest, "Сервопривод указан не верно")
	}
	srvPosition := shared.ServoPosition(uint32(srv))

	msg := server.clientMQTT.StorageServoInfo[uuid][srvPosition.String()]
	return c.JSON(fiber.Map{
		"uuid":          uuid,
		"servo":         srvPosition.String(),
		"angle":         msg.Angle,
		"duty":          msg.Duty,
		"lock":          msg.Lock,
		"move":          msg.Move,
		"chPWD":         msg.ChPWD,
		"connectionPin": msg.ConnectionPin,
	})
}

func (server *Server) handServoInfoAll(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	msgMap := fiber.Map{}
	msgMap["uuid"] = uuid
	storage := server.clientMQTT.StorageServoInfo[uuid]
	for key, val := range storage {
		msgMap[key] = map[string]any{
			"angle":         val.Angle,
			"duty":          val.Duty,
			"lock":          val.Lock,
			"move":          val.Move,
			"chPWD":         val.ChPWD,
			"connectionPin": val.ConnectionPin,
		}
	}
	return c.JSON(msgMap)
}
