package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/xwvvvvwx/builder/actions"
	"github.com/xwvvvvwx/builder/context"
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

				imageID := actions.BuildImage(context.Config().DockerfileDirectory)
				actions.StartBackgroundContainer(
					imageID, context.ContainerName(), context.Config().Volumes)

			},
		},
		{
			Name:  "exec",
			Usage: "execute a single command inside the build environment",
			Action: func(c *cli.Context) {

				command := append([]string{c.Args().First()}, c.Args().Tail()...)
				actions.ExecuteCommand(context.ContainerName(), command...)

			},
		},
		{
			Name:  "attach",
			Usage: "Attach to the project build environment. Will bring envionment up if needed",
			Action: func(c *cli.Context) {

				actions.Attach(context.ContainerName())

			},
		},
		{
			Name:  "destroy",
			Usage: "destroy the project build environment",
			Action: func(c *cli.Context) {

				actions.Destroy(context.ContainerName())

			},
		},
		{
			Name:  "clean",
			Usage: "destroy and rebuild the project build environment",
			Action: func(c *cli.Context) {

				actions.Destroy(context.ContainerName())
				imageID := actions.BuildImage(context.Config().DockerfileDirectory)
				actions.StartBackgroundContainer(
					imageID, context.ContainerName(), context.Config().Volumes)

			},
		},
	}

	app.Run(os.Args)
}
