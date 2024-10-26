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
	if cfg.JWT.PrivateKey == nil {
		log.Warn("JWT Keys are not specified. Generate random")
		cfg.GenerateFakeJWT()
		private, public, err := cfg.JWT.Export()
		if err != nil {
			log.Panic("Export keys fails", log.Err(err))
		}
		log.Info("Generated JWT Keys", log.Str("private", string(private)), log.Str("public", string(public)))
	}
	app := application.NewServerCLI(Version, cfg)
	if err := app.Run(os.Args); err != nil {
		log.Fatal("App fail", log.Err(err))
	}
}
