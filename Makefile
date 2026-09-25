postgres:
	docker run --name sba-postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:16-alpine

createdb:
	docker exec -it sba-postgres createdb --username=postgres --owner=postgres sba_db

dropdb:
	docker exec -it sba-postgres dropdb sba_db

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/sba_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/sba_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb postgres migratedown migrateup
