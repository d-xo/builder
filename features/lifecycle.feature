Feature: Lifecycle

    Scenario: Users can bring up and destroy build environments
        Given I use a fixture named "lifecycle"
        When I run `builder up`
        And I run `docker ps`
        Then the output from "builder up" should contain "Started background container with name: lifecycle-integration-test"
        And the output from "docker ps" should contain "lifecycle-integration-test"
        When I run `builder destroy`
        And I run `docker ps`
        Then the output from "builder destroy" should contain "Destroyed container with name: lifecycle-integration-test"
        And the output from "docker ps" should not contain "lifecycle-integration-test"
