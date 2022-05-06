run:
	APP_ENV=local go run main.go

build:
	go build .

unit-test:
	go test ./... -short -timeout 10s

mockgen:
	~/go/bin/mockgen -destination=internal/record/mocks/mock_repository.go -package mocks getir-assignment/internal/record Repository
	~/go/bin/mockgen -destination=internal/record/mocks/mock_service.go -package mocks getir-assignment/internal/record RecordService
	~/go/bin/mockgen -destination=internal/in_memory/mocks/mock_repository.go -package mocks getir-assignment/internal/in_memory Repository
	~/go/bin/mockgen -destination=internal/in_memory/mocks/mock_service.go -package mocks getir-assignment/internal/in_memory MemoryService

db-test:
	 go test ./internal/record/repository_test.go ./internal/record/mongodb_repository_test.go -v
 