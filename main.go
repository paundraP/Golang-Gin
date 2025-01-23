package main

import (
	"rest-api-go/internal/config"
)

func main() {
	// testing
	app := config.NewApp()
	app.Run()
}
