build: ./cmd/main.go
	rm -rf ./bin
	go build -o ./bin/server ./cmd/main.go