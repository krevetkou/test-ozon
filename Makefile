init:
	go run github.com/99designs/gqlgen init

generate:
	go run github.com/99designs/gqlgen && go run ./app/models/model_tags/model_tags.go

run:
	go run main.go

test:
	go test -v ./...

migrate_up:
	migrate -path "./app/infra/db/migration" -database "postgresql://postgres:356597@localhost:5432/postgres?sslmode=disable" up

migrate_down:
	migrate -path "./app/infra/db/migration" -database "postgresql://postgres:356597@localhost:5432/postgres?sslmode=disable" down

migrate_fix:
	migrate -path "./app/infra/db/migration" -database "postgresql://postgres:356597@localhost:5432/postgres?sslmode=disable" force 130620244
	migrate -path "./app/infra/db/migration" -database "postgresql://postgres:356597@localhost:5432/postgres?sslmode=disable" force 13062024

