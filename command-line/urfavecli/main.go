package main

import (
	"fmt"
	"os"

	"github.com/micro/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "Language for the greeting",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "greet",
			Aliases: []string{"g"},
			Usage:   "greet command",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "times, t",
					Value: 1,
					Usage: "repeat n times",
				},
			},
			Action: func(c *cli.Context) {
				for i := 0; i < c.Int("times"); i++ {
					fmt.Println("Hello World")
				}
			},
		},
	}

	app.Run(os.Args)
}
