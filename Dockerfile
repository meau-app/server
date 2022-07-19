FROM golang:1.17-bullseye

COPY cmd .
COPY internal .

ENV MEAU_HOST=0.0.0.0
ENV MEAU_PORT=80

RUN go mod tidy
RUN go build -o meau.out ./comd/main.go

EXPOSE 80 8080

ENTRYPOINT [ "meau.out" ]