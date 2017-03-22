package main

import (
	"io"
	"log"
	"os"
)

var (
	infoLogger		*log.Logger
	warningLogger	*log.Logger
	//errorLogger		*log.Logger
)

func initLogger(
	infoHandler		io.Writer,
	warningHandler	io.Writer) {
	
	infoLogger = log.New(infoHandler,
		"[INFO] ",
		log.Ldate | log.Ltime)

	warningLogger = log.New(warningHandler,
		"[WARNING] ",
		log.Ldate | log.Ltime)
}

func FatalLogger(err error) {
	log.SetPrefix("[ERROR] ")
	log.SetFlags(0)
	log.Fatal(err)
}

func ErrorOutput(err error) {
	warningLogger.Println(err)
}

func LogOutput(log string) {
	infoLogger.Println(log)
}

func InitLogger() {
	initLogger(os.Stdout, os.Stdout)
}
