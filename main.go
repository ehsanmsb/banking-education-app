package main

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/ehsanmsb/banking-education-app/internal/app"
	"github.com/ehsanmsb/banking-education-app/internal/logger"
)

// I am going to implement Hexagonal Architecture (adaptors and ports)
func main() {
	//log.Println("Starting application...")
	startText := figure.NewFigure("Banking App", "", true)
	startText.Print()
	logger.Info("Starting application")
	app.StartServer()
}
