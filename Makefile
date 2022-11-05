createdb:
	createdb --username=userPostgres --owner=userPostgres go_finance

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=senhaPostgres -d postgres:14-alpine 

migrateup:
	migrate -path db/migration -database "postgres://userPostgres:senhaPostgres@localhost:5432/go_finance?sslmode=disable" -verbose up

migrationdrop:
	migrate -path db/migration -database "postgres://userPostgres:senhaPostgres@localhost:5432/go_finance?sslmode=disable" -verbose down

sqlcgenerate:
	docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createdb postgres dropdb migrateup migrationdrop test server sqlcgenerate