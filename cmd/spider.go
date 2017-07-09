package cmd

import (
	"fmt"

	"github.com/gladmo/wechat/chat/parse/lengtoo"
	"github.com/urfave/cli"
)

var Spider = cli.Command{
	Name:        "spider",
	Aliases:     []string{"d"},
	Usage:       "Spider Command",
	Description: "Spider frame",
	Action:      runSpider,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "lengtoo",
			Usage: "Crawls lengtoo",
		},
	},
}

func runSpider(ctx *cli.Context) error {

	args := ctx.Args()
	if len(args) == 1 {
		switch args[0] {
		case "lengtoo":
			lengtoo.Spider()
		default:
			fmt.Println("Spider not defined!")
		}
	} else {
		fmt.Println("Too many parameters!")
	}

	return nil
}
