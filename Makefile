build:
	go build -o cron-parser main.go

test:
	go test ./... -cover