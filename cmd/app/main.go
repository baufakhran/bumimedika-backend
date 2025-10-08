package main

import (
	"bumimedika-backend/internal/app"
	"bumimedika-backend/internal/app/config"
)

func main() {
	cfg := config.LoadConfig()
	app.StartApp(cfg)
}
