package server

import (
	"Bionic-Web-Control/proto/shared"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (server *Server) handStrainGaugeByFingerId(c *fiber.Ctx) error {
	uuid, err := c.ParamsInt("uuid")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	fingerId := c.Params("finger_id")
	fng, fun := shared.Finger_value[fingerId]
	if !fun {
		return fiber.NewError(fiber.StatusBadRequest, "Палец указан не правильно")
	}
	fngPosition := shared.Finger(uint32(fng))

	msg, err := server.clientMQTT.StrainGaugeByFingerId(uuid, fngPosition)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// TODO дописать работу муахахахах
	return c.JSON(fiber.Map{
		"finger":        "",
		"pressure":      12,
		"connectionPin": "",
	})
}

func (server *Server) handStrainGaugeFingerAll(c *fiber.Ctx) error {
	uuid, err := c.ParamsInt("uuid")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// TODO сделать цикл который буде прогоняться через все пальцы
	msg, err := server.clientMQTT.StrainGaugeFingerAll(uuid)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	fmt.Println(msg)

	return c.JSON(fiber.Map{
		"uuid":   uuid,
		"xAccel": msg.XAccel,
		"yAccel": msg.YAccel,
		"zAccel": msg.ZAccel,
	})
}
