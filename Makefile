build:
	go build -o api cmd/main.go

test:
	go test -v ./...
