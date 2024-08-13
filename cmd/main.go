package main

import (
	"github.com/mirjalilova/auth-service/config"
	"github.com/mirjalilova/auth-service/pkg/app"
)

func main() {
	cfg := config.Load()
	
	app.Run(&cfg)
}