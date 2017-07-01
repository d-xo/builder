#! /usr/bin/env bash

gofmt -s -w ./**/*.go && golint && go test ./**/ && go install
