CC = go

SRC = cmd/main.go

all:
	$(CC) build -o meau.out $(SRC)