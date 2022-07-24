CC = go

SRC = cmd/main.go

all:
	$(CC) build -o meau.out $(SRC)


docker-build:
	docker build -t meau-server .

docker-run: docker-build
	docker run --publish 8080:8080 meau-server

docker-tag:
	docker tag meau-server gcr.io/unb-adote/meau-server

docker-push: docker-build docker-tag
	docker push gcr.io/unb-adote/meau-server