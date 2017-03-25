# Workspace

`workspace` is a tool to help manage isolated per project development environments

You can think of it as an extremely simplified version of vagrant that works only with docker containers.

## Usage

**Config**

Configuration is read from a `.workspace.json` file in the current working directory.

The following options can be set:

```
{
    "dockerfile": "<PATH_TO_DOCKERFILE_DIRECTORY>"
    "volumes": {
        "<HOST>":"<GUEST>"
    }
}
```

**Commands**

- `workspace up`: builds the environment specified in the .workspace.json
- `workspace attach`: spawns a new shell in the environment
- `workspace destroy`: destroys the environment
- `workspace reset`: destroy, rebuild, attach
