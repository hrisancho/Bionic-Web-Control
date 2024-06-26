package server

// @schemes http
func (server Server) SetupRoutes() {

	// TODO не работает, надо доделать
	server.App.Static("/", "../frontend/src/pages/Home/")
	//server.App.Static("/", "../frontend/src/pages/New1/")

	api := server.App.Group("/api")
	apiV1 := api.Group("/v1")
	apiHand := apiV1.Group("/hand")

	// TODO сделать проверку ограничений, посмотреть на сколько правильность работы при не валидируемых данных:
	//func IsValidUUID(uuid string) bool {
	//	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	//	return r.MatchString(uuid)
	//}

	// TODO стоит ли проверять входные запросы на полноту входящих данных
	// TODO При выгрузке данных удалить лишние print в терминал
	//
	// TODO если нет определённого значения стоит ли возвращать определённый статус код, если что сейчас возвращается nil
	// TODO если есть subscribe, то стоит ли делать unsubscribe ?
	// TODO сделать проверку валидацию поступающих uuid, в самом конце
	apiHand.Put("/:uuid/subscribe-to-uuid", server.handInitSubscribeToUUID)

	apiHand.Get("/:uuid/monitoring/imu/raw-data", server.handImuRawData)
	apiHand.Get("/:uuid/monitoring/imu/processed-data", server.handImuProcData)

	apiHand.Get("/:uuid/monitoring/strain-gauge/finger-id/:finger_id", server.handStrainGaugeByFingerId)
	apiHand.Get("/:uuid/monitoring/strain-gauge/all-finger", server.handStrainGaugeFingerAll)

	apiHand.Get("/:uuid/monitoring/servo/info/servo-id/:servo_id", server.handServoInfoByServoId)
	apiHand.Get("/:uuid/monitoring/servo/info/all-servo", server.handServoInfoAll)

	// TODO не доделано,3 что ниже
	apiHand.Get("/:uuid/monitoring/potentiometer/finger-id/:finger_id/position-id/:position_id", server.handPotentiometerByID)
	apiHand.Get("/:uuid/monitoring/potentiometer/finger-id/:finger_id", server.handPotentiometerByFingerId)
	apiHand.Get("/:uuid/monitoring/potentiometer/all-potentiometer", server.handPotentiometerAll)

	// Все команды ниже являются работоспособными и могут использоваться

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
