package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/xwvvvvwx/builder/actions"
	"github.com/xwvvvvwx/builder/state"
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

				imageID := actions.BuildImage(state.Config().DockerfileDirectory)
				actions.StartBackgroundContainer(
					imageID, state.ContainerName(), state.Config().Volumes)

			},
		},
		{
			Name:  "exec",
			Usage: "execute a single command inside the build environment",
			Action: func(c *cli.Context) {

				command := append([]string{c.Args().First()}, c.Args().Tail()...)
				actions.ExecuteCommand(state.ContainerName(), command...)

			},
		},
		{
			Name:  "run",
			Usage: "exec `alias` defined in .workspace.json",
			Action: func(c *cli.Context) {

				aliasName := c.Args().First()
				actions.ExecuteCommand(state.ContainerName(), state.CommandFromAlias(aliasName))

			},
		},
		{
			Name:  "attach",
			Usage: "Attach to the project build environment.",
			Action: func(c *cli.Context) {

				actions.Attach(state.ContainerName())

			},
		},
		{
			Name:  "destroy",
			Usage: "destroy the project build environment",
			Action: func(c *cli.Context) {

				actions.Destroy(state.ContainerName())

			},
		},
		{
			Name:  "clean",
			Usage: "destroy and rebuild the project build environment",
			Action: func(c *cli.Context) {

				actions.Destroy(state.ContainerName())
				imageID := actions.BuildImage(state.Config().DockerfileDirectory)
				actions.StartBackgroundContainer(
					imageID, state.ContainerName(), state.Config().Volumes)

			},
		},
		{
			Name:  "build",
			Usage: "executes the build alias",
			Action: func(c *cli.Context) {

				actions.ExecuteCommand(state.ContainerName(), state.CommandFromAlias("build"))

			},
		},
		{
			Name:  "start",
			Usage: "executes the start alias",
			Action: func(c *cli.Context) {

				actions.ExecuteCommand(state.ContainerName(), state.CommandFromAlias("start"))

			},
		},
		{
			Name:  "verify",
			Usage: "executes the verify alias",
			Action: func(c *cli.Context) {

				actions.ExecuteCommand(state.ContainerName(), state.CommandFromAlias("verify"))

			},
		},
		{
			Name:  "package",
			Usage: "executes the package alias",
			Action: func(c *cli.Context) {

				actions.ExecuteCommand(state.ContainerName(), state.CommandFromAlias("package"))

			},
		},
	}

	app.Run(os.Args)
}
