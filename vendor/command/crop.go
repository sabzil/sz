package command

import (
	"flag"
	"fmt"
	"github.com/mitchellh/cli"
	"strings"
)

type CropCommand struct {
	Ui cli.Ui
}

func (c *CropCommand) Help() string {
	helpText := `
	Usage: sz crop -src="./foo.jpg" -x=0 -y=0 -width=100 -height=20
	`
	return strings.TrimSpace(helpText)
}

func (c *CropCommand) Synopsis() string {
	return "이미지를 (startx, starty) 에서 부터 width와 height만큼 잘라줍니다"
}

func (c *CropCommand) Run(args []string) int {
	//fmt.Println("crop run")
	var src, x, y, width, height string
	cmdFlags := flag.NewFlagSet("crop", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	cmdFlags.StringVar(&src, "src", "", "")
	cmdFlags.StringVar(&x, "x", "", "")
	cmdFlags.StringVar(&y, "y", "", "")
	cmdFlags.StringVar(&width, "width", "", "")
	cmdFlags.StringVar(&height, "height", "", "")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if src == "" {
		c.Ui.Error("원본 사진의 경로를 입력하세요")
		c.Ui.Error("")
		c.Ui.Error(c.Help())
		return 1
	}

	if x == "" {
		c.Ui.Error("시작점 x를 입력하세요.")
		c.Ui.Error("")
		c.Ui.Error(c.Help())
		return 1
	}

	if y == "" {
		c.Ui.Error("시작점 y를 입력하세요.")
		c.Ui.Error("")
		c.Ui.Error(c.Help())
		return 1
	}

	if width == "" {
		c.Ui.Error("자르기 위한 넓이 width를 입력하세요.")
		c.Ui.Error("")
		c.Ui.Error(c.Help())
		return 1
	}

	if height == "" {
		c.Ui.Error("자르기 위한 높이 height를 입력하세요.")
		c.Ui.Error("")
		c.Ui.Error(c.Help())
		return 1
	}

	c.crop(src, x, y, width, height)

	return 0
}

func (c *CropCommand) crop(src, x, y, width, height string) {
	fmt.Println(src)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(width)
	fmt.Println(height)
}
