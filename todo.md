# TODO (prio sorted)

- use docker api rather than wrapping command line
- proper error messages
- default to using bare dockerfile in root
- documentation
- investigate docker shared namespace (for tooling integration)
- integration tests inside a vm
- ci
- container name can probs just be the dockerfile name (more readable than a hash)

# Error cases:

- docker is not installed
- docker deamon is not running
- no dockerfile / builder.json
- up: container is already running
- up: container exists but it not yet running
- attach / exec: container is not running
- destroy: container is not present
