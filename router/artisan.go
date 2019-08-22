package router

import (
	"flag"
	"log"
)

var tasks = make(map[string]func())

func Task() bool {
	var taskName string
	var showHeplp bool
	var server bool
	flag.BoolVar(&server, "s", true, "web server default action")
	flag.BoolVar(&showHeplp, "h", false, "usage info")
	flag.StringVar(&taskName, "t", "null", "task name")
	flag.Parse()

	if showHeplp {
		flag.Usage()
		return true
	}

	if taskName == "null" {
		return false
	}

	if task, exists := tasks[taskName]; exists {
		task()
		return true
	} else {
		log.Panicf("task %v not exists", taskName)
	}

	return true
}

func RegistTask(name string, callback func()) {
	tasks[name] = callback
}

func init() {
	RegistTask("demo", func() {
		log.Println("artisan demo,let`s do something!!")
	})
}
