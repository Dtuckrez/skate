package main

import (
	"os"
	"skate/cmd/api"
	"skate/cmd/worker"
)

func main() {
	var startMode = os.Getenv("START_MODE")
	if startMode == "api" {
		api.Start()
	} else if startMode == "worker" {
		worker.Start()
	} else {
		os.Exit(1)
	}
}
