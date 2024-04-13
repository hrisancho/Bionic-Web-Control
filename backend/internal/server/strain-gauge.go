package server

import (
	"Bionic-Web-Control/proto/shared"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (server *Server) handStrainGaugeByFingerId(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	fingerId := c.Params("finger_id")
	fng, fun := shared.Finger_value[fingerId]
	if !fun {
		return fiber.NewError(fiber.StatusBadRequest, "Палец указан не правильно")
	}
	fngPosition := shared.Finger(uint32(fng))

	msg := server.clientMQTT.StorageStrainGauge[uuid][fngPosition.String()]
	return c.JSON(fiber.Map{
		"uuid":          uuid,
		"finger":        fngPosition.String(),
		"pressure":      msg.Pressure,
		"connectionPin": msg.ConnectionPin,
	})
}

func (server *Server) handStrainGaugeFingerAll(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	msg := server.clientMQTT.StorageStrainGauge[uuid]
	fmt.Println("!!!")
	fmt.Println(msg)
	fmt.Println("!!!")
	// TODO доделать это через цикл
	for key, value := range msg {
		fmt.Println(key, value)
	}
	return nil
}
