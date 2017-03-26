package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/xwvvvvwx/builder/context"
	"github.com/xwvvvvwx/builder/docker"
)

func main() {
	app := cli.NewApp()

	app.Name = "builder"
	app.Usage = "build the machine that builds the machine"

	app.Commands = []cli.Command{
		{
			Name:  "up",
			Usage: "spin up the project build envionment",
			Action: func(c *cli.Context) {

				imageID := docker.BuildImage(context.Config().DockerfileDirectory)
				docker.StartBackgroundContainer(imageID, context.ContainerName(), context.Config().Volumes)

			},
		},
		{
			Name:  "attach",
			Usage: "Attach to the project build environment. Will bring envionment up if needed",
			Action: func(c *cli.Context) {

				docker.Attach(context.ContainerName())

			},
		},
		{
			Name:  "destroy",
			Usage: "destroy the project build environment",
			Action: func(c *cli.Context) {

				docker.Destroy(context.ContainerName())

			},
		},
		{
			Name:  "clean",
			Usage: "destroy and rebuild the current environment",
			Action: func(c *cli.Context) {

				docker.Destroy(context.ContainerName())
				imageID := docker.BuildImage(context.Config().DockerfileDirectory)
				docker.StartBackgroundContainer(imageID, context.ContainerName(), context.Config().Volumes)

			},
		},
	}

	app.Run(os.Args)
}
