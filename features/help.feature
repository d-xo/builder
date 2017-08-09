Feature: Help

    Scenario Outline: Users can show the help page
        When I run `<command>`
        Then the exit status should be 0
        And the output should contain "NAME:"
        And the output should contain "USAGE:"
        And the output should contain "VERSION:"
        And the output should contain "COMMANDS:"

        Examples:
        | command        |
        | builder        |
        | builder --help |
