install:
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@go install github.com/vektra/mockery/v2@latest
	@GO111MODULE=on go mod vendor
start:
	@go run cmd/*.go
build:
	@GOOS=linux GOARCH=amd64 go build -v cmd/*.go
clean:
	@rm -rf main
test:
	@go test -v -race ./...
swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init --parseDependency --parseDepth 1 --dir ./cmd
