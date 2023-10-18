DB_URL=postgresql://admin:password@localhost:5432/agency_banking?sslmode=disable
run:
	go run cmd/main.go
up:
	docker compose up -d

migrateup:
	migrate -path migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path  migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir migration -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...
server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go simplebank/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go simplebank/worker TaskDistributor