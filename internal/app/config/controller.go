package config

import "bumimedika-backend/internal/controllers"

type Controller struct {
	AuthController     controllers.IAuthController
	ProductController  controllers.IProductController
	CustomerController controllers.ICustomerController
}

func InitController(usecase *Usecase) *Controller {
	return &Controller{
		AuthController:     controllers.InitAuthController(usecase.AuthUsecase),
		ProductController:  controllers.InitProductController(usecase.ProductUsecase),
		CustomerController: controllers.InitCustomerController(usecase.CustomerUsecase),
	}
}
