run:
	go run ./cmd/api --port=8080

gen:
	go generate ./...

test:
	go test ./...

vet:
	go vet ./...

fmt:
	gofmt -s -w ./
	goimports -w ./