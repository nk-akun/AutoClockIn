package main

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	logFile := "/root/go/src/github.com/nk-akun/AutoClockIn/clock_in.log"
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil && os.IsNotExist(err) {
		f, _ = os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE, 0755)
	}

	log.Out = f
}

func main() {
	var action string
	flag.StringVar(&action, "a", "", "要执行的动作,report/leave")
	flag.Parse()

	if action == "report" {
		DoReport()
	} else if action == "leave" {
		Doleave()
	}
}
