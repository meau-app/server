FROM golang:1.19-bullseye

WORKDIR /app

COPY go.mod        ./
COPY go.sum        ./
COPY firebase.json ./

COPY cmd      ./cmd
COPY internal ./internal

ENV MEAU_HOST=0.0.0.0
ENV MEAU_PORT=80
ENV MEAU_AUTHENTICATE=false
ARG MEAU_VERSION

ENV GO111MODULE=on

RUN go clean -modcache
RUN go mod download
RUN go build -o ./meau.out ./cmd/main.go

EXPOSE 80 8080

ENTRYPOINT [ "./meau.out" ]