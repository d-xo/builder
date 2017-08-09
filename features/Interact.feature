Feature: Interact

    Scenario: Users can execute single commands in the build environment using exec
        Given I use a fixture named "interact"
        When I run `builder up`
        And I run `builder exec 'cat /interact.txt'`
        Then the exit status should be 0
        And the output from "builder exec 'cat /interact.txt'" should contain exactly "inside interact container"
