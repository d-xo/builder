gofmt -s -w **/*.go && gometalinter && go test ./**/ && go install
