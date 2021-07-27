test:
	@go test -v -cover ./...
build:
	@go build -o mqtt-api main.go
run:
	@PORT=8081 go run main.go
docker-build:
	@docker build -t mqtt-api:latest .
docker-run:
	@docker run -p 8081:8081 --rm mqtt-api:latest
