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

        # create file in running build env
        When I run `builder up`
        And I run `docker ps`
        And I run `builder exec "touch -c /hi.txt"`
        Then the exit status should be 0
        And the output from "docker ps" should contain "lifecycle-integration-test"

        # if file is not present after running clean then we can assume that the environment is clean
        When I run `builder clean`
        And I run `docker ps`
        And I run `builder exec 'stat -f /hi.txt'`
        And the output from "builder exec 'stat -f /hi.txt'" should contain exactly "stat: can't read file system information for '/hi.txt': No such file or directory"
        And the output from "docker ps" should contain "lifecycle-integration-test"
