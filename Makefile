CC = go

SRC = cmd/main.go

all:
	$(CC) build -o meau.out $(SRC)


docker-build:
	docker build -t meau-server .

docker-run: docker-build
	docker run -t meau-server
