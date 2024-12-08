package goatlogger

type LogLevel string

const (
	Info  LogLevel = "INFO"
	Debug LogLevel = "DEBUG"
	Error LogLevel = "ERROR"
	Panic LogLevel = "PANIC"
)
