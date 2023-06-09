package main

import (
	"github.com/daniial79/Banking-API/src/app"
	"github.com/daniial79/Banking-API/src/logger"
)

func main() {
	logger.Info("starting application...")
	app.Start()
}
