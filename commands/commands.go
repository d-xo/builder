// Package commands implements the UI.
// Pass state into actions to modify the surroundings
package commands

import (
	"github.com/urfave/cli"
	"github.com/xwvvvvwx/builder/actions"
	"github.com/xwvvvvwx/builder/data"
)

// Up starts a background container from the Dockerfile specified in the config
func Up(c *cli.Context) {
	imageID := actions.BuildImage(data.Config().DockerfileDirectory)
	actions.StartBackgroundContainer(
		imageID, data.ContainerName(), data.Config().Volumes)
}

// Exec executes a single command in the build environment
func Exec(c *cli.Context) {
	command := append([]string{c.Args().First()}, c.Args().Tail()...)
	actions.ExecuteCommand(data.ContainerName(), command...)
}

// Run executes the specified alias
func Run(c *cli.Context) {
	aliasName := c.Args().First()
	executeAlias(aliasName)
}

// Attach spawns a bash shell in the build environment
func Attach(c *cli.Context) {
	actions.Attach(data.ContainerName())
}

// Destroy destroys the build environment
func Destroy(c *cli.Context) {
	actions.Destroy(data.ContainerName())
}

// Clean resets the build environment to the state specified in the Dockerfile
func Clean(c *cli.Context) {
	Destroy(c)
	Up(c)
}

// Build executes the "build" alias
func Build(c *cli.Context) {
	executeAlias("build")
}

// Start executes the "start" alias
func Start(c *cli.Context) {
	executeAlias("start")
}

// Verify executes the "verify" alias
func Verify(c *cli.Context) {
	executeAlias("verify")
}

// Package executes the "package" alias
func Package(c *cli.Context) {
	executeAlias("package")
}

func executeAlias(aliasName string) {
	actions.ExecuteCommand(data.ContainerName(), data.CommandFromAlias(aliasName))
}
