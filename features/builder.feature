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
        Then the features should all pass
        And the output should contain "Started background container with name:"
