install:
	@go get -u github.com/pressly/goose/cmd/goose
	@GO111MODULE=on go mod vendor
start:
	@go run cmd/*.go
build:
	@GOOS=linux GOARCH=amd64 go build -v cmd/*.go
clean:
	@rm -rf main
test:
	@go test -v -race ./...
