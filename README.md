# builder

`builder` is a tool to help manage dockerized project build environments.

You can think of it as something like vagrant mixed with npm.

Define your build environment in a Dockerfile and then define custom user commands to be executed in
that build environment.

The environment persists in the background between command executions to allow for incremental compilation.

## Motivation

- Each project has an isolated and reproducible build environment versioned alongside the code
- Easily work on different projects with incompatible build environments on the same machine
- Developers new to the project can be up and running with a functional build environment in seconds

## Usage

**Config**

Configuration is read from a `.workspace.json` file in the project root

The following options can be set:

```
{
    "dockerfileDirectory": "<PATH_TO_DOCKERFILE_DIRECTORY>",
    "volumes": {
        "<HOST>":"<GUEST>"
    },
    "commands": {
        "name1": "command",
        "name2": "command",
        "name3": "command"
    }
}
```

**Core Commands**

- `builder up`: spins up the environment specified in the .workspace.json
- `builder exec`: executes a single command in the build environment
- `builder attach`: spawns a new bash shell in the build environment
- `builder destroy`: destroys the build environment
- `builder clean`: reset the environment to the state specified in the .workspace.json

**User Defined Commands**

Users can define command aliases in the .workspace.json. These commands can be accessed via `builder run commandName`.

The following aliases can be accessed without the `run` keyword:

- `builder build`
- `builder verify`
- `builder package`
