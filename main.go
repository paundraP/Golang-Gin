package main

import (
	"realtime-score/internal/config"
)

func main() {
	app := config.NewApp()
	app.Run()
}
