package mq

import (
	"fmt"

	"github.com/urfave/cli"
)

type MqPeek struct {
	cli.Command
}

func NewMqPeek() *MqPeek {
	mqPeek := &MqPeek{
		Command: cli.Command{
			Name:      "peek",
			Usage:     "do the doo",
			UsageText: "doo - does the dooing",
			ArgsUsage: "[image] [args]",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: test ", c.Args().First())
				return nil
			},
		},
	}

	return mqPeek
}

func (r MqPeek) GetCmd() cli.Command {
	return r.Command
}