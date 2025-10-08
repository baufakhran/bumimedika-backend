package app

import (
	"bumimedika-backend/internal/app/config"
	"fmt"
)

func StartApp(cfg *config.Config) {
	// init repository
	repository := config.InitRepository(cfg)
	// init usecase
	usecase := config.InitUsecase(repository, cfg)
	// init controller
	controller := config.InitController(usecase)
	// init routes
	routes := initRoutes(*controller)
	// register routes
	routes.registerRoutes(cfg, repository.Rds)

	if err := router.Run(fmt.Sprintf(":%s", cfg.AppPort)); err != nil {
		panic(err)
	}
}
