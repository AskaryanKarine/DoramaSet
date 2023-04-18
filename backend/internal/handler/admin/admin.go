package admin

import "DoramaSet/internal/interfaces/controller"

type Admin struct {
	dc controller.IDoramaController
	sc controller.IStaffController
}

func New(dc controller.IDoramaController, sc controller.IStaffController) *Admin {
	return &Admin{
		dc: dc,
		sc: sc,
	}
}

// func (a *Admin) CreateDorama(token string) error {
//
// }
