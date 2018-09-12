package main

import (
	"log"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./log/test.log",
		MaxSize:    10, // megabytes
		MaxBackups: 10,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})

	log.Print("test log")
}
