package server

import (
	"Bionic-Web-Control/proto/potentiometer"
	"Bionic-Web-Control/proto/shared"
	"github.com/gofiber/fiber/v2"
)

func (server *Server) handPotentiometerByID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	// работа с пальцами
	fingerId := c.Params("finger_id")
	_, fun := shared.Finger_value[fingerId]
	if !fun {
		return fiber.NewError(fiber.StatusBadRequest, "Палец указан не правильно")
	}

	positionId := c.Params("position_id")
	_, fun = potentiometer.Position_value[positionId]
	if !fun {
		return fiber.NewError(fiber.StatusBadRequest, "Позиция потенциометра указана не верно")
	}
	if server.clientMQTT.StoragePotentiometerAngle[uuid][fingerId] == nil {
		return c.JSON(fiber.Map{"uuid": uuid})
	}
	storage := server.clientMQTT.StoragePotentiometerAngle[uuid][fingerId][positionId]
	return c.JSON(fiber.Map{
		"uuid":          uuid,
		"finger_id":     fingerId,
		"position_id":   positionId,
		"angle":         storage.Angle,
		"connectionPin": storage.ConnectionPin,
	})
}

func (server *Server) handPotentiometerByFingerId(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	// работа с пальцами
	fingerId := c.Params("finger_id")
	_, fun := shared.Finger_value[fingerId]
	if !fun {
		return fiber.NewError(fiber.StatusBadRequest, "Палец указан не правильно")
	}
	msgMap := fiber.Map{}
	msgMap["uuid"] = uuid
	storage := server.clientMQTT.StoragePotentiometerAngle[uuid][fingerId]
	for key, val := range storage {
		msgMap[key] = map[string]any{
			"angle":         val.Angle,
			"connectionPin": val.ConnectionPin,
		}
	}
	return c.JSON(msgMap)
}

func (server *Server) handPotentiometerAll(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	msgMap := fiber.Map{}
	msgMap["uuid"] = uuid

	storage := server.clientMQTT.StoragePotentiometerAngle[uuid]
	//Уровень пальцев
	for key_fing, val_fing := range storage {
		// Уровень позиции
		anonMap := make(map[string]any)
		for key_pos, val_pos := range val_fing {
			anonMap[key_pos] = map[string]any{
				"angle":         val_pos.Angle,
				"connectionPin": val_pos.ConnectionPin,
			}
		}
		msgMap[key_fing] = anonMap
	}
	return c.JSON(msgMap)
}
