lint:
	@golangci-lint run --deadline=5m

test:
	go test -count=1 -v -p 1 $(shell go list ./...| grep -v /vendor/)

benchtest:
	go test -bench=.
