# builder

`builder` is a tool to help manage isolated per project build environments


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
    ""
}
```

**Core Commands**

- `builder up`: builds the environment specified in the .workspace.json
- `builder attach`: spawns a new shell in the environment
- `builder destroy`: destroys the environment
- `builder clean`: destroy, rebuild, attach
