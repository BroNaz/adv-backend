.PHONY: build, test, migrate_up, migrate_test_up, migrate_down, migrate_test_down

build:
		go build -v ./cmd/apiserver 

test: 
		go test -v -race -timeout 30s ./...

migrate_up: 
		migrate -path migrations -database "postgresql://postgres:12@localhost/restapi_dev?sslmode=disable" up 

migrate_test_up:
		migrate -path migrations -database "postgresql://postgres:12@localhost/restapi_test?sslmode=disable" up

migrate_down:
		migrate -path migrations -database "postgresql://postgres:12@localhost/restapi_dev?sslmode=disable" down

migrate_test_down:
		migrate -path migrations -database "postgresql://postgres:12@localhost/restapi_test?sslmode=disable" down

.DEFAULT_GOAL := build