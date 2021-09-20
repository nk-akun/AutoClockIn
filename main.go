package main

import (
	"flag"
	"fmt"
)

func main() {
	var action string
	flag.StringVar(&action, "a", "", "要执行的动作,report/leave")
	flag.Parse()

	fmt.Println(action)

	if action == "report" {
		DoReport()
	} else if action == "leave" {
		Doleave()
	}
}
