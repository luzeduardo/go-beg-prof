build:
	go build -o api cmd/main.go

test:
	go test -v -short ./...

coverage:
	go test ./... -cover

report:
	go tool cover -html=coverage.out -o cover.xhtml
	
