package middleware

import (
	"os"
	"time"
)

var LogCh = make(chan string, 100)

func AsyncLogger() {
	os.MkdirAll("storage/logs", 0755)
	f, _ := os.OpenFile("storage/logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	for msg := range LogCh {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		f.WriteString("[" + timestamp + "] " + msg + "\n")
	}
}
