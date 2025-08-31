package main

import "logger"

func Process(logger logger.Logger) {
	logger.Info("Processing started")
}

func main() {
	log := logger.ConsoleLogger{}
	Process(log)
}
