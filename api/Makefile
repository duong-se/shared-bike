install:
	@go install github.com/githubnemo/CompileDaemon@latest
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@go install github.com/vektra/mockery/v2@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@GO111MODULE=on go mod vendor
start:
	@go run *.go
build:
	@GOOS=linux GOARCH=amd64 go build  -o cmd/main -v *.go
clean:
	@rm -rf cmd/main
test:
	@go test -v -race ./...
swagger:
	@swag init --parseDependency --parseDepth 1
migrations:
	@goose -dir ./sql/migrations mysql $DB_CONNECTION_STRING up
seeders:
	@goose -dir ./sql/seeders mysql $DB_CONNECTION_STRING up
