package main

import (
	"gopkg.in/urfave/cli.v1"
	"fmt"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Count up or down"
	app.Commands = []cli.Command{
		{
			Name:      "up",
			ShortName: "u",
			Usage:     "Count up",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "stop,s",
					Usage: "value to count up to",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("stop")
				if start <= 0 {
					fmt.Println("Stop cannot be negative.")
				}
				for i := 0; i <= start; i++ {
					time.Sleep(time.Second)
					fmt.Println(i)
				}
				return nil
			},
		},
		{
			Name:      "down",
			ShortName: "d",
			Usage:     "Count down",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "start,s",
					Usage: "start counting down from",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("start")
				if start < 0 {
					fmt.Println("start cannot be negative.")
				}
				for i := start; i >= 0; i-- {
					time.Sleep(time.Second)
					fmt.Println(i)
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
