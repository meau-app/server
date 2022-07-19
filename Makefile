CC = go

SRC = cmd/main.go

all:
	$(CC) build -o meau.out $(SRC)


docker-build:
	docker build -t meau-server .

docker-run: docker-build
	docker run --publish 8080:8080 meau-server
