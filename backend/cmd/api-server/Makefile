export DB_NAME=api
export DB_USER=api-server
export DB_HOST=localhost
export DB_PASSWORD=password
export DB_PORT=5432

export AUTH_TOKEN=mytoken
export PORT=7000

build:
	go build

run: build
	./api-server

ensure-deps:
	rm Gopkg.lock; dep ensure
