package application

import (
	"errors"
	"fmt"

	"github.com/sashaaro/gophkeeper/internal/config"
	"github.com/sashaaro/gophkeeper/internal/server"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func NewServerCLI(version string, cfg *config.Server) *cli.App {
	return &cli.App{
		Name:    "GophKeeper server",
		Version: version,
		Usage:   "Run and use",
		Commands: []*cli.Command{
			{
				Name: "serve",
				Action: func(context *cli.Context) error {
					grpcServer := server.NewGRPCServer(cfg.Listen)
					if err := grpcServer.Serve(); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
						return err
					}
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Printf("GophKeeper %s\n", c.App.Version)
			return nil
		},
		DefaultCommand: "serve",
	}
}
