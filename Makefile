build:
	GO111MODULE=on go build -o cron-parser main.go

test:
	go test ./... -cover