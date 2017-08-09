Feature: Lifecycle

    Scenario: Users can build docker containers
        Given I use a fixture named "lifecycle"
        When I run `builder up`
        And I run `docker ps`
        Then the output from "builder up" should contain "Started background container with name: lifecycle-integration-test"
        And the output from "docker ps" should contain "lifecycle-integration-test"
