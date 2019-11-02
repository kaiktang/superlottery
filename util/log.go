package util

import (
	"github.com/op/go-logging"
	"os"
)

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func init() {
	logFile, err := os.OpenFile("./log.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	backend1 := logging.NewLogBackend(logFile, "", 0)

	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.INFO, "")

	logging.SetBackend(backend1Leveled, backend1Formatter)
}

func GetLogger(module string) *logging.Logger {
	return logging.MustGetLogger(module)
}

