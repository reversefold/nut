package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "nut"
	app.Usage = "the development environment, containerized"
	// define nut subcommands
	app.Commands = []cli.Command{
		{
			Name:  "status",
			Usage: "gives status of the dev env",
			Action: func(c *cli.Context) {
				status()
			},
		},
		{
			Name:  "build",
			Usage: "build project in a container",
			Action: func(c *cli.Context) {
				build()
			},
		},
		{
			Name:  "run",
			Usage: "run project in a container",
			Action: func(c *cli.Context) {
				run()
			},
		},
	}

	app.Run(os.Args)
}