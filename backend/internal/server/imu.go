package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (server *Server) handImuRawData(c *fiber.Ctx) error {
	uuid, err := c.ParamsInt("uuid")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	msg, err := server.clientMQTT.ImuRawData(uuid)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	fmt.Println(msg)

	return c.JSON(fiber.Map{
		"uuid":          uuid,
		"number":        msg.Number,
		"xAccel":        msg.XAccel,
		"yAccel":        msg.YAccel,
		"zAccel":        msg.ZAccel,
		"xAngle":        msg.XAngle,
		"yAngle":        msg.YAngle,
		"zAngle":        msg.ZAngle,
		"connectionPin": msg.ConnectionPin,
	})
}

func (server *Server) handImuProcData(c *fiber.Ctx) error {
	uuid, err := c.ParamsInt("uuid")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	msg, err := server.clientMQTT.ImuProcData(uuid)
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
