Feature: Exec

    Scenario: Users can execute single commands in the build environment using exec
        Given I use a fixture named "exec"
        When I run `builder up`
        And I run `builder exec 'cat /exec.txt'`
        Then the exit status should be 0
        And the output from "builder exec 'cat /exec.txt'" should contain exactly "inside exec container"
