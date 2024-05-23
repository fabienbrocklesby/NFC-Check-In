package logger

import (
	"log"
	"os"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
	Debug(msg string)
}

type logger struct {
	log *log.Logger
}

func NewLogger() Logger {
	return &logger{
		log: log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *logger) Info(msg string) {
	l.log.Println("INFO: " + msg)
}

func (l *logger) Error(msg string) {
	l.log.Println("ERROR: " + msg)
}

func (l *logger) Debug(msg string) {
	l.log.Println("DEBUG: " + msg)
}
