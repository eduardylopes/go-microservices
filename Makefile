MICROSERVICES:=authentication-service broker-service listener-service logger-service mailer-service front-end
FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp
AUTH_BINARY=authApp
LOGGER_BINARY=loggerApp
MAILER_BINARY=mailerApp
LISTENER_BINARY=listenerApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_auth build_logger build_mailer build_listener build_front
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"

## build: build all binaries
build: build_broker build_auth build_logger build_mailer build_listener build_front

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ./broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ./bin/${BROKER_BINARY} ./cmd/api
	@echo "Done!"

## build_auth: builds the auth binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd ./authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ./bin/${AUTH_BINARY} ./cmd/api
	@echo "Done!"

## build_logger: builds the auth binary as a linux executable
build_logger:
	@echo "Building logger binary..."
	cd ./logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ./bin/${LOGGER_BINARY} ./cmd/api
	@echo "Done!"

## build_mailer: builds the auth binary as a linux executable
build_mailer:
	@echo "Building mailer binary..."
	cd ./mailer-service && env GOOS=linux CGO_ENABLED=0 go build -o ./bin/${MAILER_BINARY} ./cmd/api
	@echo "Done!"

## build_mailer: builds the auth binary as a linux executable
build_listener:
	@echo "Building listener binary..."
	cd ./listener-service && env GOOS=linux CGO_ENABLED=0 go build -o ./bin/${LISTENER_BINARY}
	@echo "Done!"

## start: starts the front end
start: build_front
	@echo "Starting front end"
	cd ./front-end && ./bin/${FRONT_END_BINARY}

## build_front: builds the frone end binary
build_front:
	@echo "Building front-end binary..."
	cd ./front-end && env CGO_ENABLED=0 go build -o ./bin/${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"

## migrateup: runs all the migrations
migrateup:
	migrate -path ./authentication-service/db/migration -database "postgresql://postgres:password@localhost:5432/users?sslmode=disable" -verbose up

## migratedown: downs all the migrations
migratedown:
	migrate -path ./authentication-service/db/migration -database "postgresql://postgres:password@localhost:5432/users?sslmode=disable" -verbose down $(steps)

## migratecreate: creates migrations files up and down
migratecreate:
	migrate create -ext sql -dir ./authentication-service/db/migration -seq $(name).sql

## upload_images: upload images to the docker
install_dependencies:
	cd ./authentication-service && go mod tidy
	cd ./broker-service && go mod tidy
	cd ./listener-service && go mod tidy
	cd ./logger-service && go mod tidy
	cd ./mailer-service && go mod tidy
	cd ./front-end && go mod tidy