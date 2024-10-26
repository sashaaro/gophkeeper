package application

import (
	"fmt"

	"github.com/sashaaro/gophkeeper/internal/client"
	"github.com/sashaaro/gophkeeper/internal/config"
	"github.com/sashaaro/gophkeeper/internal/log"
	"github.com/urfave/cli/v2"
)

func NewClientCLI(version string, cfg *config.Client) *cli.App {
	return &cli.App{
		Name:    "GophKeeper client",
		Version: version,
		Usage:   "say a greeting",
		Commands: []*cli.Command{
			{
				Name: "ping",
				Action: func(ctx *cli.Context) error {
					grpcClient := client.NewGRPCClient(cfg.ServerAddr, client.WithoutTLS())
					defer func() {
						if err := grpcClient.Close(); err != nil {
							log.Error("close grpc fail", log.Err(err))
						}
					}()
					return grpcClient.Ping(ctx.Context)
				},
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Printf("GophKeeper %s\n", c.App.Version)
			return nil
		},
	}
}
