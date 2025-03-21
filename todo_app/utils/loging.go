package utils

import (
	"io"
	"log"
	"os"
)

var logfile *os.File

func LoggingSettings(logFile string) {
	var err error
	logfile, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func CloseLogFile() {
	if logfile != nil {
		logfile.Close()
	}
}
