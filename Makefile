run:
	go run main.go

test:
	godog run --format=pretty features/*.feature