package main

import (
	"os"

	"command"
	"github.com/mitchellh/cli"
)

var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	Commands = map[string]cli.CommandFactory{
		"crop": func() (cli.Command, error) {
			return &command.CropCommand{
				Ui: ui,
			}, nil
		},
	}
}
