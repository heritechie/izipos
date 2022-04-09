postgres:
	 docker run --name postgres12 --network izipos-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root izipos

dropdb:
	docker exec -it postgres12 dropdb --username=root  izipos

migrateup:
	migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5432/izipos?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5432/izipos?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run cmd/server/main.go

mock:
	mockgen -package mockdb -destination internal/db/mock/store.go github.com/heritechie/izipos/api/internal/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock