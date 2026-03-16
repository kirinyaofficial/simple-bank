postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secrete -d postgres:17.5-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank

createmigrate: 
	migrate create -ext sql -dir db/migration -seq ${name}

migrateup:
	migrate -path db/migration/ -database "postgresql://root:secrete@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration/ -database "postgresql://root:secrete@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration/ -database "postgresql://root:secrete@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration/ -database "postgresql://root:secrete@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

psql:
	docker exec -it postgres17 psql -U root simple_bank

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go github.com/kirinyaofficial/simple-bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc psql test server mock createmigrate