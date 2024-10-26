package main

import (
	"os"

	"github.com/sashaaro/gophkeeper/internal/application"
	"github.com/sashaaro/gophkeeper/internal/config"
	"github.com/sashaaro/gophkeeper/internal/log"
)

var Version = "0.0.0"

func main() {
	cfg := config.NewServer()
	app := application.NewServerCLI(Version, cfg)
	if err := app.Run(os.Args); err != nil {
		log.Fatal("App fail", log.Err(err))
	}
}
