include .env

migrateup:
	migrate -path=db/migration -database "postgresql://${PG_USERNAME}:${PG_PASSWORD}@${PG_HOST}:${PG_PORT}/${PG_DATABASE}?sslmode=disable" up

migratedown:
	migrate -path=db/migration -database "postgresql://${PG_USERNAME}:${PG_PASSWORD}@${PG_HOST}:${PG_PORT}/${PG_DATABASE}?sslmode=disable" down

install:
	go install github.com/codegangsta/gin@latest&&
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest &&
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest &&
	go install github.com/golang/mock/mockgen@v1.6.0

sqlc:
	sqlc generate
	
test:
	go test -v -cover ./...

mock:
	go generate ./...

.PHONY: migrateup migratedown install sqlc test mock