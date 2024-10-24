package main

import (
	"fmt"
	"os"

	"github.com/sashaaro/gophkeeper/internal/log"
	"github.com/urfave/cli/v2"
)

var Version = "0.0.0"

func main() {
	app := &cli.App{
		Name:    "GophKeeper client",
		Version: Version,
		Usage:   "say a greeting",
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
