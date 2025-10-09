FROM golang:1.25 AS build

WORKDIR /src
COPY go.mod go.sum /src/
RUN go mod download && go mod verify

COPY . /src/

RUN CGO_ENABLED=0 GOOS=linux go build -o /server /src/cmd/server/main.go

FROM alpine:latest AS release

WORKDIR /
RUN apk add --no-cache mysql-client mariadb-connector-c

COPY --from=build /server /server

EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/server"]