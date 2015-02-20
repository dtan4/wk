package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "wk"
	app.Version = Version
	app.Usage = ""
	app.Author = "Daisuke Fujita"
	app.Email = "dtanshi45@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
