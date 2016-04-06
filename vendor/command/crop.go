package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"strings"
)

type CropCommand struct {
	Ui cli.Ui
}

func (c *CropCommand) Help() string {
	helpText := `
	Usage: sz crop -src="./foo.jpg" -startx=0 -starty=0 -width=100 -height=20
	`
	return strings.TrimSpace(helpText)
}

func (c *CropCommand) Synopsis() string {
	return "이미지를 (startx, starty) 에서 부터 width와 height만큼 잘라줍니다"
}

func (c *CropCommand) Run(args []string) int {
	//fmt.Println("crop run")
	if len(args) == 0 {
		fmt.Println(c.Help())
		return 0
	}
	return 0
}
