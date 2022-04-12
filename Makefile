run:
	go run main.go

test:
	godog run --format=pretty

ginkgo-test:
	ginkgo -r