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
	var grpcClient *client.GRPCClient
	return &cli.App{
		Name:    "GophKeeper client",
		Version: version,
		Usage:   "say a greeting",
		Commands: []*cli.Command{
			{
				Name: "ui",
				Action: func(c *cli.Context) error {
					uiApp := ui.NewUIApp(client.NewClient(cfg))
					uiApp.Init()
					return uiApp.Run()
				},
			},
			{
				Name: "ping",
				Action: func(ctx *cli.Context) error {
					if err := grpcClient.Ping(ctx.Context); err != nil {
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
					if err := grpcClient.Register(ctx.Context, login, password); err != nil {
						return err
					}
					log.Info("Registered")
					return nil
				},
			},
		},
		Before: func(c *cli.Context) error {
			grpcClient = client.NewGRPCClient(cfg.ServerAddr, client.WithoutTLS())
			fmt.Printf("GophKeeper %s\n", c.App.Version)
			return nil
		},
		After: func(c *cli.Context) error {
			if err := grpcClient.Close(); err != nil {
				log.Error("close grpc fail", log.Err(err))
			}
			return nil
		},
	}
}
