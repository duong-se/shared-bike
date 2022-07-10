# Wait for DB starting
wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"
# Run migration
goose -dir ./sql/migrations mysql $DB_CONNECTION_STRING up
# Run seeders
goose -dir ./sql/seeders mysql $DB_CONNECTION_STRING up
# Run service and monitor changes
$GOPATH/bin/CompileDaemon --build="go build -o main main.go"  --command=./main
