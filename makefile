build:
	dep ensure
	go test ./...
	go build -o bin/go-evolve ./main.go