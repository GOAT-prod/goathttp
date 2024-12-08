package goatlogger

import (
	"encoding/json"
	"fmt"
	"time"
)

var (
	moscowTimeLocation, _ = time.LoadLocation("Europe/Moscow")
)

func printLog(level LogLevel, app, tag, msg string) {
	info := logInfo{
		Application: app,
		Level:       level,
		Time:        time.Now().In(moscowTimeLocation).Format(time.RFC3339),
		Tag:         tag,
		Msg:         msg,
	}

	infoBytes, _ := json.Marshal(info)
	fmt.Println(string(infoBytes))
}
