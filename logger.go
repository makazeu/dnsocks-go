package main

import (
	"io"
	"log"
	"os"
)

var (
	comLogger		*log.Logger
	infoLogger		*log.Logger
	warningLogger	*log.Logger
	//errorLogger		*log.Logger
)

func initLogger(
	comHandler		io.Writer,
	infoHandler		io.Writer,
	warningHandler	io.Writer) {
	
	comLogger = log.New(comHandler,
		"", 0)

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

func CommonOutput(text string) {
	comLogger.Println(text)
}

func ErrorOutput(err error) {
	warningLogger.Println(err)
}

func LogOutput(log string) {
	infoLogger.Println(log)
}

func InitLogger() {
	initLogger(
		os.Stdout,	// Common Logger
		os.Stdout,  // Debug Logger
		os.Stdout)	// Error Logger
}
