package command

import (
	"core"
	"flag"
	"github.com/mitchellh/cli"
	"strings"
)

type ResizeCommand struct {
	Ui cli.Ui
}

func (r *ResizeCommand) Help() string {
	helpText := `
	Usage: sz resize -source="./foo.jpg" -width=10 -height=30 -filter="MitchellNetravali"

	filter: NearestNeighbor, Box, Linear, Hermite, MitchellNetravali, CatmullRom, BSpline, Gaussian, Lanczos, Hann, Hamming, Blackman, Bartlett, Welch, Cosine
	`

	return strings.TrimSpace(helpText)
}

func (r *ResizeCommand) Synopsis() string {
	return "이미지를 확대/축소 합니다."
}

func (r *ResizeCommand) Run(args []string) int {
	var source, width, height, filter string
	cmdFlags := flag.NewFlagSet("resize", flag.ContinueOnError)
	cmdFlags.Usage = func() { r.Ui.Output(r.Help()) }
	cmdFlags.StringVar(&source, "source", "", "")
	cmdFlags.StringVar(&width, "width", "", "")
	cmdFlags.StringVar(&height, "height", "", "")
	cmdFlags.StringVar(&filter, "filter", "", "")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if source == "" {
		r.Ui.Error("원본 사진의 경로를 입력하세요.")
		r.Ui.Error("")
		r.Ui.Error(r.Help())
		return 1
	}

	if width == "" {
		r.Ui.Error("확대/축소할 이미지의 넓이를 입력하세요.")
		r.Ui.Error("")
		r.Ui.Error(r.Help())
		return 1
	}

	if height == "" {
		r.Ui.Error("확대/축소할 이미지의 높이를 입력하세요.")
		r.Ui.Error("")
		r.Ui.Error(r.Help())
		return 1
	}

	if filter == "" {
		r.Ui.Error("확대/축소할 필터를 선택해주세요.")
		r.Ui.Error("")
		r.Ui.Error(r.Help())
		return 1
	}

	core.Resize(source, width, height, filter)
	return 0
}
