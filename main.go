package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/xwvvvvwx/builder/commands"
)

func main() {
	app := cli.NewApp()

	app.Name = "builder"
	app.Usage = "build the machine that builds the machine"

	app.Commands = []cli.Command{
		{
			Name:   "up",
			Usage:  "spin up the project build envionment",
			Action: commands.Up,
		},
		{
			Name:   "exec",
			Usage:  "execute a single command inside the build environment",
			Action: commands.Exec,
		},
		{
			Name:   "run",
			Usage:  "exec `alias` defined in .builder.json",
			Action: commands.Run,
		},
		{
			Name:   "attach",
			Usage:  "Attach to the project build environment.",
			Action: commands.Attach,
		},
		{
			Name:   "destroy",
			Usage:  "destroy the project build environment",
			Action: commands.Destroy,
		},
		{
			Name:   "clean",
			Usage:  "destroy and rebuild the project build environment",
			Action: commands.Clean,
		},
		{
			Name:   "build",
			Usage:  "executes the build alias",
			Action: commands.Build,
		},
		{
			Name:   "start",
			Usage:  "executes the start alias",
			Action: commands.Start,
		},
		{
			Name:   "verify",
			Usage:  "executes the verify alias",
			Action: commands.Verify,
		},
		{
			Name:   "package",
			Usage:  "executes the package alias",
			Action: commands.Package,
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
