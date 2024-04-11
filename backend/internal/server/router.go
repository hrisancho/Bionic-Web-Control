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

	//
	//// Company endpoints
	//apiCompany := apiV1.Group("/company")
	//apiCompany.Post("/", middlewareLoginRequired, middlewareSuperAdminRequired, server.companyCreate)
	//apiCompany.Put("/", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.companyUpdate)
	//apiCompany.Delete("/:id", middlewareLoginRequired, middlewareSuperAdminRequired, server.companyDelete)
	//apiCompany.Post("/list", middlewareLoginRequired, middlewareSuperAdminRequired, server.companyList)
	//apiCompany.Get("/all", middlewareLoginRequired, middlewareSuperAdminRequired, server.companyAll)
	//apiCompany.Get("/uuid/:uuid", server.companyGetByUUID)
	//apiCompany.Get("/:id", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.companyGetById)
	//apiCompany.Get("/:id/users", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.companyGetAllUsers)
	//
	//apiCompanySearch := apiCompany.Group("/search")
	//apiCompanySearch.Get("/name", middlewareLoginRequired, middlewareSuperAdminRequired, server.companyGetAllLikeName)
	//
	//// User endpoints
	//apiUser := apiV1.Group("/user")
	//apiUser.Post("/signin", server.userSignIn)
	//apiUser.Post("/signup", server.userSignUp)
	//apiUser.Get("/current", middlewareLoginRequired, server.userCurrent)
	//apiUser.Get("/not-accepted", middlewareLoginRequired, middlewareAdminRequired, server.userGetNotAccepted)
	//apiUser.Post("/list", middlewareLoginRequired, middlewareSuperAdminRequired, server.userList)
	//apiUser.Post("/list/company/:companyId", middlewareLoginRequired, middlewareAdminRequired, server.userListByCompany)
	//apiUser.Get("/all", middlewareLoginRequired, middlewareSuperAdminRequired, server.userAll)
	//apiUser.Get("/all/company/:companyId", middlewareLoginRequired, middlewareAdminRequired, server.userAllByCompany)
	//apiUser.Get("/:id", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.userGet)
	//apiUser.Post("/", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.userCreate)
	//apiUser.Put("/", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.userUpdate)
	//apiUser.Post("/:id/photo", middlewareLoginRequired, server.userSetPhoto)
	//apiUser.Get("/:id/photo", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.userGetPhoto)
	//apiUser.Delete("/:id", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.userDelete)
	//apiUser.Put("/:id/accept", middlewareLoginRequired, middlewareAdminRequired, server.userAccept)
	//apiUser.Put("/:id/decline", middlewareLoginRequired, middlewareAdminRequired, server.userDecline)
	//apiUser.Put("/:id/user-group/:userGroupId", middlewareLoginRequired, middlewareAdminRequired, server.userAddToUserGroup)
	//
	//apiUserTg := apiUser.Group("/tg")
	//apiUserTg.Get("/bind-link", middlewareLoginRequired, server.userTgCreateBindToken)
	//apiUserTg.Post("/bind", server.userTgBind)
	//apiUserTg.Get("/created-at/:id", server.userGetCreatedAtByTgId)
	//
	//apiUserAccess := apiUser.Group("/access")
	//// Get user node access list
	//apiUserAccess.Get("/:id/nodes", middlewareLoginRequired, middlewareSuperAdminRequired, server.userGetAllNodes)
	//// Add user access to node
	//apiUserAccess.Post("/node", middlewareLoginRequired, middlewareAdminRequired, server.userAddAccessNode)
	//// Add demo user access to node
	//apiUserAccess.Post("/node/demo", middlewareLoginRequired, middlewareAdminRequired, server.userAddAccessNodeDemo)
	//// Delete access to node
	//apiUserAccess.Delete("/:userId/node/:nodeId", middlewareLoginRequired, middlewareAdminRequired, server.userDeleteAccessNode)
	//
	//// User-group endpoints
	//apiUserGroup := apiV1.Group("/user-group")
	//apiUserGroup.Post("/", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.userGroupCreate)
	//apiUserGroup.Put("/", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.userGroupUpdate)
	//apiUserGroup.Delete("/:id", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.userGroupDelete)
	//apiUserGroup.Post("/list", middlewareLoginRequired, middlewareSuperAdminRequired, server.userGroupList)
	//apiUserGroup.Post("/list/company/:id", middlewareLoginRequired, middlewareAdminRequired, server.userGroupListByCompany)
	//apiUserGroup.Get("/all", middlewareLoginRequired, middlewareSuperAdminRequired, server.userGroupAll)
	//apiUserGroup.Get("/all/company/:id", middlewareLoginRequired, middlewareAdminRequired, server.userGroupAllByCompany)
	//
	//// User-moving-history endpoints
	//apiUserMovingHistory := apiV1.Group("/user-moving-history")
	//apiUserMovingHistory.Post("/list", middlewareLoginRequired, middlewareAdminRequired, server.userMovingHistoryList)
	//
	//// User-action-history endpoints
	//apiUserActionHistory := apiV1.Group("/user-action-history")
	//apiUserActionHistory.Post("/list", middlewareLoginRequired, middlewareAdminRequired, server.userActionHistoryList)
	//apiUserActionHistory.Post("/list/company/:id", middlewareLoginRequired, middlewareAdminOrSecurityRequired, server.userActionHistoryListByCompany)
	//
	//// Node endpoints
	//apiNode := apiV1.Group("/node")
	//apiNode.Get("/controller-updates-ws", webSocketUpgrader, middlewareLoginRequiredWebSocket, middlewareAdminOrSuperAdminRequired, server.nodeAllControllersUpdatesWS())
	//apiNode.Post("/", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.nodeCreate)
	//apiNode.Put("/", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.nodeUpdate)
	//apiNode.Delete("/:id", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.nodeDelete)
	//apiNode.Post("/list", middlewareLoginRequired, middlewareSuperAdminRequired, server.nodeList)
	//apiNode.Post("/list/company/:id", middlewareLoginRequired, middlewareAdminRequired, server.nodeListByCompany)
	//apiNode.Get("/all", middlewareLoginRequired, middlewareSuperAdminRequired, server.nodeAll)
	//apiNode.Get("/all/company/:companyId", middlewareLoginRequired, middlewareAdminRequired, server.nodeGetAllByCompany)
	//apiNode.Get("/:id", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.nodeGetById)
	//apiNode.Get("/:id/controllers", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.nodeGetAllControllers)
	//apiNode.Get("/:id/users", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.nodeGetAllUsers)
	//
	//apiNodeBind := apiNode.Group("/bind")
	//apiNodeBind.Post("/controller", middlewareLoginRequired, middlewareAdminRequired, server.nodeBindController)
	//apiNodeBind.Delete("/:nodeId/controller/:controllerId", middlewareLoginRequired, middlewareAdminRequired, server.nodeUnbindController)
	//
	//// Controller endpoints
	//apiController := apiV1.Group("/controller")
	//apiController.Post("/", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.controllerCreate)
	//apiController.Put("/", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.controllerUpdate)
	//apiController.Delete("/:id", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.controllerDelete)
	//apiController.Get("/scanner", middlewareLoginRequired, middlewareAdminRequired, server.controllerScanner)
	//apiController.Post("/list", middlewareLoginRequired, middlewareSuperAdminRequired, server.controllerList)
	//apiController.Post("/list/company/:id", middlewareLoginRequired, middlewareAdminRequired, server.controllerListByCompany)
	//apiController.Get("/all", middlewareLoginRequired, middlewareSuperAdminRequired, server.controllerAll)
	//apiController.Get("/all/company/:id", middlewareLoginRequired, middlewareAdminRequired, server.controllerAllByCompany)
	//apiController.Get("/all/company/:id/online", middlewareLoginRequired, middlewareAdminRequired, server.controllerAllByCompanyOnline)
	//apiController.Get("/all/node/:nodeId", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.controllerGetAllByNode)
	//apiController.Get("/:id", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.controllerGetById)
	//apiController.Get("/uuid/:uuid", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.controllerGetByUUID)
	//apiController.Put("/:id/enable", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.controllerEnable)
	//apiController.Put("/:id/disable", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.controllerDisable)
	//apiController.Get("/:firmware_uuid/ota", server.controllerOtaRead)
	//apiController.Post("/ota", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.controllerOtaRequest)
	//apiControllerControl := apiController.Group("/control")
	//apiControllerControl.Post("/door", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.controllerControlDoor)
	////apiControllerControl.Post("/photo", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.controllerControlPhoto)
	//apiControllerConfig := apiController.Group("/config")
	//apiControllerConfig.Post("/get", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.controllerConfigGet)
	//apiControllerConfig.Post("/set", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.controllerConfigSet)
	//apiControllerConfig.Post("/set/name", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.controllerSetName)
	//apiControllerConfig.Post("/reset", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.controllerConfigReset)
	//
	//// Controller notification endpoints
	//apiControllerNotification := apiV1.Group("/controller-notification")
	//apiControllerNotification.Post("/list", middlewareLoginRequired, middlewareAdminOrSuperAdminRequired, server.controllerNotificationList)
	//
	//// Notification description list
	//apiNotificationDescription := apiV1.Group("/notification-description")
	//apiNotificationDescription.Get("/group-map", middlewareLoginRequired, server.notificationDescriptionGroupMap)
	//apiNotificationDescription.Get("/list", middlewareLoginRequired, server.notificationDescriptionList)
	//
	//// Notification subscription endpoints
	//apiNotificationSubscription := apiV1.Group("/notification-subscription")
	//apiNotificationSubscription.Post("/", middlewareLoginRequired, middlewareAdminRequired, server.notificationSubscriptionCreate)
	//apiNotificationSubscription.Post("/delete", middlewareLoginRequired, middlewareAdminRequired, server.notificationSubscriptionDelete)
	//apiNotificationSubscription.Post("/list", middlewareLoginRequired, middlewareAdminRequired, server.notificationSubscriptionList)
}
