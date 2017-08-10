Feature: Volumes

    Scenario: Volumes can be bound from the host to the guest
        Given I use a fixture named "volumes"
        And I run `builder up`
        When I run `builder exec 'cat /vol1/vol1.txt'`
        And I run `builder exec 'cat /vol2/vol2.txt'`
        Then the output from "builder exec 'cat /vol1/vol1.txt'" should contain exactly "this file is in vol1"
        And the output from "builder exec 'cat /vol2/vol2.txt'" should contain exactly "this file is in vol2"

