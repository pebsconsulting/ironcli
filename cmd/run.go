package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

type Run struct {
	cli.Command
}

func NewRun() *Run {
	run := &Run{
		Command: cli.Command{
			Name:      "run",
			Usage:     "do the doo",
			UsageText: "doo - does the dooing",
			ArgsUsage: "[image] [args]",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
		},
	}

	return run
}

func (r Run) GetCmd() cli.Command {
	return r.Command
}