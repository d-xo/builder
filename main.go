package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/xwvvvvwx/builder/actions"
	"github.com/xwvvvvwx/builder/transforms"
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

				imageID := actions.BuildImage(transforms.Config().DockerfileDirectory)
				actions.StartBackgroundContainer(
					imageID, transforms.ContainerName(), transforms.Config().Volumes)

			},
		},
		{
			Name:  "exec",
			Usage: "execute a single command inside the build environment",
			Action: func(c *cli.Context) {

				command := append([]string{c.Args().First()}, c.Args().Tail()...)
				actions.ExecuteCommand(transforms.ContainerName(), command...)

			},
		},
		{
			Name:  "run",
			Usage: "exec `alias` defined in .workspace.json",
			Action: func(c *cli.Context) {

				aliasName := c.Args().First()
				actions.ExecuteCommand(transforms.ContainerName(), transforms.CommandFromAlias(aliasName))

			},
		},
		{
			Name:  "attach",
			Usage: "Attach to the project build environment.",
			Action: func(c *cli.Context) {

				actions.Attach(transforms.ContainerName())

			},
		},
		{
			Name:  "destroy",
			Usage: "destroy the project build environment",
			Action: func(c *cli.Context) {

				actions.Destroy(transforms.ContainerName())

			},
		},
		{
			Name:  "clean",
			Usage: "destroy and rebuild the project build environment",
			Action: func(c *cli.Context) {

				actions.Destroy(transforms.ContainerName())
				imageID := actions.BuildImage(transforms.Config().DockerfileDirectory)
				actions.StartBackgroundContainer(
					imageID, transforms.ContainerName(), transforms.Config().Volumes)

			},
		},
		{
			Name:  "build",
			Usage: "executes the build alias",
			Action: func(c *cli.Context) {

				actions.ExecuteCommand(transforms.ContainerName(), transforms.CommandFromAlias("build"))

			},
		},
		{
			Name:  "start",
			Usage: "executes the start alias",
			Action: func(c *cli.Context) {

				actions.ExecuteCommand(transforms.ContainerName(), transforms.CommandFromAlias("start"))

			},
		},
		{
			Name:  "verify",
			Usage: "executes the verify alias",
			Action: func(c *cli.Context) {

				actions.ExecuteCommand(transforms.ContainerName(), transforms.CommandFromAlias("verify"))

			},
		},
		{
			Name:  "package",
			Usage: "executes the package alias",
			Action: func(c *cli.Context) {

				actions.ExecuteCommand(transforms.ContainerName(), transforms.CommandFromAlias("package"))

			},
		},
	}

	app.Run(os.Args)
}
