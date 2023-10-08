PWD=$(shell pwd)
SERVICE=eater-svc
MIGRATION_PATH=${PWD}/src/infrastructure/migrations
PROTOS_PATH=$(PWD)/src/infrastructure/protos

server:
	go run main.go

migration:
	migrate create -ext sql -dir pkg/database/migrations -seq $(table)

migrateup:
	migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable&search_path=public" -path ./pkg/database/migrations up

migratedown:
	migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable&search_path=public" -path ./pkg/database/migrations down

PWD=$(shell pwd)
PROTOS_PATH=$(PWD)/src/infrastructure/protos

clear:
	rm -rf genprotos/*
gen-eater:
	protoc \
    --go_out=./src/application/protos \
	--go_opt=paths=import \
	--go-grpc_out=./src/application/protos \
	--go-grpc_opt=paths=import \
	-I=$(PWD)/src/infrastructure/protos/eater \
	$(PWD)/src/infrastructure/protos/eater/*.proto


docker: 
	docker build --rm -t eater-svc -f ${PWD}/deploy/Dockerfile .

add-protos-submodule:
	git submodule add https://github.com/Azamjon99/services-proto.git ./src/infrastructure/protos

pull-protos-submodule:
	git submodule update --recursive --remote
	



server-docker: server
	docker build --rm -t server-app -f ./docker/server/Dockerfile .

deploy-server:
	docker run --rm -p 3030:3030 --name=app server-app

deploy-client:
	docker run --rm --network=host --name=app client-app

compose-up:
	docker-compose -f deploy/docker-compose.yml up

compose-down:
	docker-compose -f deploy/docker-compose.yml down	
network-create:
	docker network create -d bridge mcs-network


gen-restaurant-staff-proto:
	protoc \
	--go_out=./src/application/protos \
	--go_opt=paths=import \
	--go-grpc_out=./src/application/protos \
	--go-grpc_opt=paths=import \
	-I=$(PROTOS_PATH)/restaurant_staff \
	$(PROTOS_PATH)/restaurant_staff/*.proto
