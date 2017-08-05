#! /usr/bin/env ash
# shellcheck shell=dash

echo ""
echo "Format / Lint"
echo "==================================================="
gofmt -s -w ./**/*.go
golint

echo ""
echo "Install Dependencies"
echo "==================================================="
go get

echo ""
echo "Unit Tests"
echo "==================================================="
go test ./...

echo ""
echo "Integration Tests"
echo "==================================================="
go install
cucumber
