.PHONY: test db

test:
	govendor test +local

deps:
	go get -t ./...

install:
	govendor install +local +vendor

schema:
	(cd api/routes/v1/schema; go-bindata -o schema.go -pkg schema -prefix assets assets/)

queries:
	(cd db/queries; go-bindata -o queries.go -pkg queries -prefix assets assets/)

build: schema queries
	go build -o bin/gin-example .

run:
	bin/gin-example api

build-docker:
	docker build -t gin-example .

shell: build-docker
	docker run --rm -it -P \
		--name=gin-example \
		--link gin-example-db:postgres \
		-v `pwd`:/go/src/github.com/replicatedcom/gin-example \
		gin-example \
		/bin/bash

run-docker: build build-docker
	docker run --rm -P \
		--name=gin-example \
		--link gin-example-db:postgres \
		-v `pwd`:/go/src/github.com/replicatedcom/gin-example \
		gin-example

db:
	docker run -d \
		--name=gin-example-db \
		 -e POSTGRES_PASSWORD=password \
		postgres:9

db-up:
	docker run --rm \
		--link gin-example-db:postgres \
		-v `pwd`:/go/src/github.com/replicatedcom/gin-example \
		gin-example \
		goose -dir db/migrations postgres "user=postgres password=password host=postgres dbname=postgres sslmode=disable" up

db-down:
	docker run --rm \
		--link gin-example-db:postgres \
		-v `pwd`:/go/src/github.com/replicatedcom/gin-example \
		gin-example \
		goose -dir db/migrations postgres "user=postgres password=password host=postgres dbname=postgres sslmode=disable" down

alias:
	mkdir -p $(GOPATH)/src/github.com/replicatedcom
	hln `pwd` $(GOPATH)/src/github.com/replicatedcom/gin-example # OS X
