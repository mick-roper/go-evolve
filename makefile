build:
	dep ensure
	go test ./...
	go build -o bin/wall-crawl ./main.go