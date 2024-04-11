package server

import "github.com/gofiber/fiber/v2"

// @schemes http
func (server Server) SetupRoutes() {
	server.App.Get("/",
		func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})

	api := server.App.Group("/api")
	apiV1 := api.Group("/v1")
	apiHand := apiV1.Group("/hand")

	// TODO изменить названия запросов
	// /api/v1/hand/:uuid/commands/servo-go-to-angle/:servo_id
	// JSON:
	//{
	//	"angle":41.7 (float)
	//}
	apiHand.Put("/:uuid/commands/servo-go-to-angle/servo-id/:servo_id", server.handServoToAngle)

	// Без JSON
	apiHand.Put("/:uuid/commands/servo-lock/servo-id/:servo_id", server.handServoLock)

	// Без JSON
	apiHand.Put("/:uuid/commands/servo-unlock/servo-id/:servo_id", server.handServoUnlock)

	//JSON
	//{
	//    "easing":"linear",
	//    "speed":32.3,
	//    "targetAngle":33.2
	//}
	apiHand.Put("/:uuid/commands/servo-smoothly-move/servo-id/:servo_id", server.handServoSmoothlyMove)

	// Без JSON
	apiHand.Put("/:uuid/commands/move-to-target-pressure/servo-id/:servo_id/finger-id/:finger_id", server.handMoveToTargetPressure)
	//JSON
	//{
	//    "gesture":"linear",
	//    "duration":32,
	//    "permanent2":33
	//}
	apiHand.Put("/:uuid/commands/hold-gesture", server.handServoHoldGesture)
}
