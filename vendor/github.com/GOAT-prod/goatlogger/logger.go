package goatlogger

import "os"

type Logger struct {
	Application string `json:"app"`
	Tag         string `json:"tag"`
}

type logInfo struct {
	Application string   `json:"app"`
	Level       LogLevel `json:"level"`
	Time        string   `json:"time"`
	Tag         string   `json:"tag"`
	Msg         string   `json:"message"`
}

func New(app string) Logger {
	return Logger{
		Application: app,
	}
}

func (l *Logger) SetTag(tag string) {
	l.Tag = tag
}

func (l *Logger) Info(message string) {
	printLog(Info, l.Application, l.Tag, message)
}

func (l *Logger) Debug(message string) {
	printLog(Debug, l.Application, l.Tag, message)
}

func (l *Logger) Error(message string) {
	printLog(Error, l.Application, l.Tag, message)
}

func (l *Logger) Panic(message string) {
	printLog(Panic, l.Application, l.Tag, message)
	os.Exit(1)
}
