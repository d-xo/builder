Feature: Lifecycle

    Scenario: Users can bring up and destroy build environments
        Given I use a fixture named "lifecycle"

        When I run `builder up`
        And I run `docker ps`
        Then the output from "docker ps" should contain "lifecycle-integration-test"

        When I run `builder destroy`
        And I run `docker ps`
        Then the output from "docker ps" should not contain "lifecycle-integration-test"


    Scenario: Clean should destroy and rebuild the container
        Given I use a fixture named "lifecycle"

        When I run `builder up`
        And I run `docker ps`
        And I run `builder exec touch /hi.txt`
        Then the output from "docker ps" should contain "lifecycle-integration-test"

        When I run `builder clean`
        And I run `docker ps`
        And I run `builder exec touch /hi.txt`
        And I run `builder exec 'stat -f /hi.txt'`
        Then the output from "builder exec 'stat -f /hi.txt'" should contain "No such file or directory"
        And the output from "docker ps" should contain "lifecycle-integration-test"
