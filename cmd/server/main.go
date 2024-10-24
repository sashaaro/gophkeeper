package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/sashaaro/gophkeeper/internal/config"
	"github.com/sashaaro/gophkeeper/internal/log"
	"github.com/sashaaro/gophkeeper/internal/server"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

var Version = "0.0.0"

func main() {
	app := &cli.App{
		Name:    "GophKeeper server",
		Version: Version,
		Usage:   "Run and use",
		Commands: []*cli.Command{
			{
				Name: "serve",
				Action: func(context *cli.Context) error {
					cfg := config.NewServer()
					srv := server.NewGRPCServer(cfg.Listen)
					if err := srv.Serve(); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
						return err
					}
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Greetings")
			return nil
		},
		DefaultCommand: "serve",
	}

	if err := app.Run(os.Args); err != nil {
		log.Error("App fail", log.Err(err))
		os.Exit(1)
	}
}
