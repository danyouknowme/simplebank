DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

network:
	docker network create bank-network
.PHONY: network

postgres:
	docker run --name postgresql --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
.PHONY: postgres

createdb:
	docker exec -it postgresql createdb --username=root --owner=root simple_bank
.PHONY: createdb

dropdb:
	docker exec -it postgresql dropdb simple_bank
.PHONY: dropdb

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
.PHONY: migrateup

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1
.PHONY: migrateup1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
.PHONY: migratedown

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1
.PHONY: migratedown1

migration:
	migrate create -ext sql -dir db/migration -seq $(name)
.PHONY: migration

db_docs:
	dbdocs build docs/db.dbml
.PHONY: db_docs

db_schema:
	dbml2sql --postgres -o docs/schema.sql docs/db.dbml
.PHONY: db_schema

sqlc:
	sqlc generate
.PHONY: sqlc

test:
	go test -v -cover ./...
.PHONY: test

server:
	go run main.go
.PHONY: server

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/danyouknowme/simplebank/db/sqlc Store
.PHONY: mock

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=docs/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
	proto/*.proto
.PHONY: proto

evans:
	evans --host localhost --port 9090 -r repl
.PHONY: evans