package docker

import (
	"os"

	"github.com/urfave/cli"
	"github.com/xwvvvvwx/workspace/docker"
	"github.com/xwvvvvwx/workspace/state"
)

func main() {
	app := cli.NewApp()

	app.Name = "env"
	app.Usage = "easily work with dockerized build envionments"

	app.Commands = []cli.Command{
		{
			Name:   "up",
			Usage:  "spin up the project build envionment",
			Action: up,
		},
		{
			Name:   "attach",
			Usage:  "Attach to the project build environment. Will bring envionment up if needed",
			Action: attach,
		},
		{
			Name:   "destroy",
			Usage:  "destroy the project build environment",
			Action: destroy,
		},
		{
			Name:   "clean",
			Usage:  "destroy and rebuild the current environment",
			Action: clean,
		},
	}

	app.Run(os.Args)
}

func up(context *cli.Context) {
	config := state.Config()
	imageID := docker.BuildImage(config)
	docker.StartBackgroundContainer(imageID, config.Volumes)
}

func attach(context *cli.Context) {
	docker.Docker("exec", "-i", "-t", state.ContainerName(), "/bin/bash")
}

func destroy(context *cli.Context) {
	docker.Docker("rm", "--force", state.ContainerName())
}

func clean(context *cli.Context) {
	destroy(context)
	up(context)
}
