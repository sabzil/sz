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
		"rotate": func() (cli.Command, error) {
			return &command.RotateCommand{
				Ui: ui,
			}, nil
		},
		"resize": func() (cli.Command, error) {
			return &command.ResizeCommand{
				Ui: ui,
			}, nil
		},
	}
}
