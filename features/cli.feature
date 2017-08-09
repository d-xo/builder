Feature: cli

    Scenario: Running builder with no commands should print help output
        When I run `builder`
        Then the features should all pass
        And the output should contain "NAME:"
        And the output should contain "USAGE:"
        And the output should contain "VERSION:"
        And the output should contain "COMMANDS:"


