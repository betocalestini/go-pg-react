# createdb:
# 	createdb --username=postgres --owner=postgres go_finance

# postgres:
# 	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine 

migrateup:
	migrate -path db/migration -database "postgres://betocalestini:cpi10@host.docker.internal:5432/udemy?sslmode=disable" -verbose up

migrationdrop:
	migrate -path db/migration -database "postgres://betocalestini:cpi10@host.docker.internal:5432/udemy?sslmode=disable" -verbose down

sqlcgenerate:
	docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createdb postgres dropdb migrateup migrationdrop test server