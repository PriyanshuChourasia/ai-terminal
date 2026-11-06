package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type LOGLEVEL int

const (
	LevelDebug LOGLEVEL = iota
	LevelInfo
	LevelError
	LevelWarn
	LevelFatal
	LevelUnique
)

const (
	colorReset  = "\033[0m"
	colorDebug  = "\033[36m"
	colorInfo   = "\033[32m"
	colorError  = "\033[31m"
	colorWarn   = "\033[33m"
	colorFatal  = "\033[35m"
	colorUnique = "\033[38m"
	colorTime   = "\033[90m"
)

var globalLevel = LevelDebug
var isLoc = loadIST()

var fileOutput io.Writer

func SetFileOutput(w io.Writer) {
	fileOutput = w
}

func SetLogLevel(level LOGLEVEL) {
	globalLevel = level
}

func loadIST() *time.Location {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		return time.FixedZone("IST", 5*3600+30*60)
	}
	return loc
}

func writeLog(loglevel LOGLEVEL, colorCode, prefix, format string, args ...any) {
	log.Printf("LogLevel %d\n", loglevel)
	if globalLevel > loglevel {
		log.Printf("Please return")
		return
	}

	timestamp := time.Now().In(isLoc).Format("2006-01-02 15:04:05.000")

	message := fmt.Sprintf(format, args...)

	fmt.Fprintf(os.Stdout, "%s%s%s %s[%s]%s %s", colorTime, timestamp, colorReset, colorCode, prefix, colorReset, message)

}

func Debug(format string, args ...any) {
	writeLog(LevelDebug, colorDebug, "DEBUG", format, args...)
}

func Info(format string, args ...any) {
	writeLog(LevelInfo, colorInfo, "INFO", format, args...)
}
