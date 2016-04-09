package command

import (
	"core"
	"flag"
	"github.com/mitchellh/cli"
	"strings"
)

type RotateCommand struct {
	Ui cli.Ui
}

func (r *RotateCommand) Help() string {
	helpText := `
	Usage: sz rotate -source="./foo.jpg" -direction="right"
	`

	return strings.TrimSpace(helpText)
}

func (r *RotateCommand) Synopsis() string {
	return "이미지를 90단위로 left, right 회전을 합니다."
}

func (r *RotateCommand) Run(args []string) int {
	var source, direction string
	cmdFlags := flag.NewFlagSet("rotate", flag.ContinueOnError)
	cmdFlags.Usage = func() { r.Ui.Output(r.Help()) }
	cmdFlags.StringVar(&source, "source", "", "")
	cmdFlags.StringVar(&direction, "direction", "", "")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if source == "" {
		r.Ui.Error("원본 사진의 경로를 입력하세요.")
		r.Ui.Error("")
		r.Ui.Error(r.Help())
		return 1
	}

	if direction == "" {
		r.Ui.Error("회전할 방향을 입력하세요.")
		r.Ui.Error("")
		r.Ui.Error(r.Help())
	}

	core.Rotate(source, direction)

	return 0
}
