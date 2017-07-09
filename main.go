package main

import (
	"os"
	"sort"

	"github.com/gladmo/wechat/cmd"
	"github.com/urfave/cli"
)

const APP_VER = "1.0.0"

func main() {
	app := cli.NewApp()

	app.Name = "Wechat spider"
	app.Usage = "Wechat spider"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.Spider,
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	app.Run(os.Args)
}
