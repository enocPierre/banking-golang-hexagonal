package main

import (
	//"log"

	"github.com/golang/Hexagonal-golangBancking/app"
	"github.com/golang/Hexagonal-golangBancking/logger"
)

func main() {
	//log.Println("Application is runing..")
	logger.Info("Starting the application...")
	app.Start()
}
