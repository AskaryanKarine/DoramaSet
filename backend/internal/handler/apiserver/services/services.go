package services

import "DoramaSet/internal/interfaces/controller"

type Services struct {
	controller.IUserController
	controller.IDoramaController
	controller.IStaffController
	controller.IEpisodeController
	controller.IListController
	controller.IPictureController
	controller.ISubscriptionController
	controller.IPointsController
}
