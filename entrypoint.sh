wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

# Run migration
goose -dir ./sql mysql $DB_CONNECTION_STRING up
$GOPATH/bin/CompileDaemon --build="go build -o main main.go"  --command=./main
