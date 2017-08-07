Feature: Aliases

    Scenario Outline: Built In
        Given I use a fixture named "aliases"
        When I run `builder up`
        And I run `builder <alias>`
        Then the output from "builder <alias>" should contain exactly "echo 'ran <alias>'\nran <alias>"

        Examples:
            | alias     |
            | build     |
            | verify    |
            | package   |
            | start     |
            | benchmark |

    Scenario: Custom
        Given I use a fixture named "aliases"
        When I run `builder up`
        And I run `builder run custom`
        Then the output from "builder run custom" should contain exactly "echo 'ran custom'\nran custom"
