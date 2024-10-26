package application

import (
	"context"
	"fmt"

	"github.com/sashaaro/gophkeeper/internal/client"
	"github.com/sashaaro/gophkeeper/internal/config"
	"github.com/sashaaro/gophkeeper/internal/log"
	"github.com/sashaaro/gophkeeper/internal/ui"
	"github.com/urfave/cli/v2"
)

type Pinger interface {
	Ping(ctx context.Context) error
}

func NewClientCLI(version string, cfg *config.Client) *cli.App {
	return &cli.App{
		Name:    "GophKeeper client",
		Version: version,
		Usage:   "say a greeting",
		Commands: []*cli.Command{
			{
				Name: "ui",
				Action: func(c *cli.Context) error {
					cl := client.NewClient(cfg)
					defer cl.Close()
					uiApp := ui.NewUIApp(cl)
					uiApp.Init()
					return uiApp.Run()
				},
			},
			{
				Name: "ping",
				Action: func(ctx *cli.Context) error {
					cl := client.NewClient(cfg)
					defer cl.Close()
					if err := cl.Ping(ctx.Context); err != nil {
						fmt.Printf("Fails. %v\n", err)
					}
					fmt.Println("PONG")
					return nil
				},
			},
			{
				Name:      "register",
				ArgsUsage: "{login} {password}", // @fixme Небезопасно оставлять пароль в истории cli
				Action: func(ctx *cli.Context) error {
					login := ctx.Args().Get(0)
					password := ctx.Args().Get(1)
					cl := client.NewClient(cfg)
					defer cl.Close()
					if err := cl.Register(ctx.Context, login, password); err != nil {
						return err
					}
					log.Info("Registered")
					return nil
				},
			},
		},
		Before: func(c *cli.Context) error {
			fmt.Printf("GophKeeper %s\n", c.App.Version)
			return nil
		},
	}
}
