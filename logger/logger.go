package logger

import "fmt"

type Logger interface {
	Info(msg string)
	Error(msg string)
}

type ConsoleLogger struct {
}

func (c ConsoleLogger) Info(msg string) {
	fmt.Println("[INFO]", msg)
}

func (c ConsoleLogger) Error(msg string) {
	fmt.Println("[ERROR]", msg)
}