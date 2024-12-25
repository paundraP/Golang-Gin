package main

import (
	"rest-api-go/internal/config"
)

func main() {
	app := config.NewApp()
	app.Run()
}
