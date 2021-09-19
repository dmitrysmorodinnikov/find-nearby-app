.PHONY: all
all: setup build lint test

assign-vars = $(if $(1),$(1),$(shell grep '$(2): ' application.yml | tail -n1| cut -d':' -f2 | cut -d' ' -f2))

APP=find_nearby_backend
APP_PACKAGES=$(shell go list ./...)
APP_EXECUTABLE="./out/$(APP)"

setup:
	export GO111MODULE=on
	if [ ! -e $(shell go env GOPATH)/bin/golangci-lint ] ; then curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.23.1 ; fi;

build:
	mkdir -p out/
	go build -o $(APP_EXECUTABLE)

run:
	go run main.go start

seed:
	go run main.go seed

install:
	go install ./...

lint:
	GO111MODULE=on golangci-lint run --disable-all \
	--enable=staticcheck --enable=unused --enable=gosimple --enable=structcheck --enable=varcheck --enable=ineffassign \
	--enable=deadcode --enable=stylecheck --enable=unconvert --enable=gofmt \
	--enable=unparam --enable=nakedret --enable=gochecknoinits --enable=depguard --enable=gocyclo --enable=misspell \
	--enable=megacheck --enable=goimports --enable=golint --enable=govet --enable=gocritic \
	--enable=scopelint --enable=rowserrcheck \
	--exclude='Using the variable on range scope \`tt\` in function literal' \
	--deadline=5m --no-config

docker.start:
	docker-compose up --remove-orphans -d

# Create a docker container or start the existing one
db.docker-start:
	docker start find-nearby-backend-dev || docker run -p 5430:5432 --name find-nearby-backend-dev -e POSTGRES_PASSWORD=postgres -d postgis/postgis

db.setup: db.create db.migrate

db.create:
	env PGPASSWORD=postgres createdb -Eutf8 -U postgres -h localhost -p 5430 find_nearby_dev

db.migrate:
	go run main.go migrate

db.rollback:
	go run main.go rollback

db.drop:
	env PGPASSWORD=postgres dropdb --if-exists -U postgres -h localhost -p 5430 find_nearby_dev

db.reset: db.docker-start db.drop db.create db.migrate

db.create-migration:
	migrate create -ext sql -dir database/migrations $(MIGRATION_NAME)

test: copy-config
	go test -p 1 -count 1 -timeout 30s -race -failfast ./...

test-cover-html:
	@echo "mode: count" > coverage-all.out
	$(foreach pkg, $(APP_PACKAGES),\
	go test -coverprofile=coverage.out -covermode=count $(pkg);\
	tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html coverage-all.out -o out/coverage.html

test-coverage:
	mkdir -p ./out/
	go test $(APP_PACKAGES) -p=1 -coverprofile ./out/cover.out
	go tool cover -func ./out/cover.out | tail -1 | awk -v thr=70 '{err=($$3 > thr) ? 0 : 1; printf("Test coverage %s is %s threshold %s%% \n", $$3, (err==0)?"above":"below",thr)} END {exit err}'

copy-config:
	cp application.yml.sample application.yml


