package server

import (
	"Bionic-Web-Control/proto/shared"
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

// TODO указать различие написание uuid'ов в одиночном запросе, мы получим его вместе со всем запросом, а если все пальцы то как ключ
func (server *Server) handStrainGaugeFingerAll(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	storage := server.clientMQTT.StorageStrainGauge[uuid]
	msgMap := fiber.Map{}
	msgMap["uuid"] = uuid
	for key, val := range storage {
		msgMap[key] = map[string]any{
			"pressure":      val.Pressure,
			"connectionPin": val.ConnectionPin,
		}
	}
	return c.JSON(msgMap)
}
