Feature: Builder

    Scenario: Running builder with no commands should print help output
        When I run `builder`
        Then the features should all pass
        And the output should contain "NAME:"
        And the output should contain "USAGE:"
        And the output should contain "VERSION:"
        And the output should contain "COMMANDS:"

    Scenario: Users can build docker containers
        When I run `builder up`
        And I run `docker ps`
        Then the features should all pass
        And the output from "builder up" should contain "Started background container with name: a1de5f24deed6d63eede585c5ccd158bf22ac830"
        And the output from "docker ps" should contain "a1de5f24deed6d63eede585c5ccd158bf22ac830"
