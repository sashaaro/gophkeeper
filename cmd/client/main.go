package main

import (
	"fmt"
	"os"

	"github.com/sashaaro/gophkeeper/internal/client"
	"github.com/sashaaro/gophkeeper/internal/config"
	"github.com/sashaaro/gophkeeper/internal/log"
	"github.com/urfave/cli/v2"
)

var Version = "v0.0.0"

func main() {
	app := &cli.App{
		Name:    "GophKeeper client",
		Version: Version,
		Usage:   "say a greeting",
		Commands: []*cli.Command{
			{
				Name: "ping",
				Action: func(ctx *cli.Context) error {
					cfg := config.NewClient()
					app := client.NewGRPCClient(cfg.ServerAddr, client.WithoutTLS())
					defer func() {
						if err := app.Stop(); err != nil {
							log.Error("close grpc fail", log.Err(err))
						}
					}()
					return app.Ping(ctx.Context)
				},
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Greetings")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Error("Fail", log.Err(err))
		os.Exit(1)
	}
}
