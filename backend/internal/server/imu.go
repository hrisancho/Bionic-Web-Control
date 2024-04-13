package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (server *Server) handImuRawData(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	msg := server.clientMQTT.StorageRawImu[uuid]
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
	uuid := c.Params("uuid")

	msg := server.clientMQTT.StorageProcessImu[uuid]
	fmt.Println(msg)

	return c.JSON(fiber.Map{
		"uuid":   uuid,
		"xAccel": msg.XPos,
		"yAccel": msg.YPos,
		"zAccel": msg.ZPos,
		"xAngle": msg.XAngle,
		"yAngle": msg.YAngle,
		"zAngle": msg.ZAngle,
	})
}
