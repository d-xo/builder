package commands

import (
	"github.com/urfave/cli"
	"github.com/xwvvvvwx/builder/actions"
	"github.com/xwvvvvwx/builder/state"
)

func Up(c *cli.Context) {
	imageID := actions.BuildImage(state.Config().DockerfileDirectory)
	actions.StartBackgroundContainer(
		imageID, state.ContainerName(), state.Config().Volumes)
}

func Exec(c *cli.Context) {
	command := append([]string{c.Args().First()}, c.Args().Tail()...)
	actions.ExecuteCommand(state.ContainerName(), command...)
}

func Run(c *cli.Context) {
	aliasName := c.Args().First()
	executeAlias(aliasName)
}

func Attach(c *cli.Context) {
	actions.Attach(state.ContainerName())
}

func Destroy(c *cli.Context) {
	actions.Destroy(state.ContainerName())
}

func Clean(c *cli.Context) {
	Destroy(c)
	Up(c)
}

func Build(c *cli.Context) {
	executeAlias("build")
}

func Start(c *cli.Context) {
	executeAlias("start")
}

func Verify(c *cli.Context) {
	executeAlias("verify")
}

func Package(c *cli.Context) {
	executeAlias("package")
}

func executeAlias(aliasName string) {
	actions.ExecuteCommand(state.ContainerName(), state.CommandFromAlias(aliasName))
}
