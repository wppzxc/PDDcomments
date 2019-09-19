package log

import (
	"log"
	"os"
)

var Logger *log.Logger
var logFile = "PDDComments.log"
var file *os.File

func init() {
	_, err := os.Stat(logFile)
	if os.IsNotExist(err) {
		os.Create(logFile)
	}
	file, err = os.OpenFile(logFile, os.O_APPEND|os.O_CREATE, 666)
	if err != nil {
		log.Fatalln("fail to create PDDComments.log file!")
	}
	Logger = log.New(file, "", log.LstdFlags | log.Lshortfile)
	Logger.SetFlags(log.LstdFlags | log.Lshortfile)
}

func Close() {
	file.Close()
}
