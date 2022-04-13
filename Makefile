run:
	go run main.go

godog-test:
	godog run --format=pretty --tags=api

ginkgo-test:
	ginkgo -r

test:
	go test ./...
