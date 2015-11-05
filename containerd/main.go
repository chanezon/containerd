package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

const (
	Version = "0.0.1"
	Usage   = `High performance conatiner daemon`
)

func main() {
	app := cli.NewApp()
	app.Name = "containerd"
	app.Version = Version
	app.Usage = Usage
	app.Authors = []cli.Author{
		Name:  "@crosbymichael",
		Email: "crosbymichael@gmail.com",
	}
	app.Commands = []cli.Command{
		DaemonCommand,
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}