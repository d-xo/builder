Feature: Aliases

    Scenario: Build
        Given I use a fixture named "aliases"
        When I run `builder up`
        And I run `builder build`
        Then the output from "builder build" should contain exactly "ran build"
