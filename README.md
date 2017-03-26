# Workspace

`workspace` is a tool to help manage isolated per project build environments

You can think of it as an extremely simplified version of vagrant that works only with docker containers.

## Usage

**Config**

Configuration is read from a `.workspace.json` file in the project root

The following options can be set:

```
{
    "dockerfileDirectory": "<PATH_TO_DOCKERFILE_DIRECTORY>"
    "volumes": {
        "<HOST>":"<GUEST>"
    }
}
```

**Commands**

- `workspace up`: builds the environment specified in the .workspace.json
- `workspace attach`: spawns a new shell in the environment
- `workspace destroy`: destroys the environment
- `workspace clean`: destroy, rebuild, attach
