run:
	go run main.go

godog-test:
	godog run --format=pretty

ginkgo-test:
	ginkgo -r

test:
	go test ./...
