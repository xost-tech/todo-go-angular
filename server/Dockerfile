FROM golang:1.22.0-alpine3.19 AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/todo-app-server

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN pwd

COPY . .

RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/todo-app-server .

FROM alpine:latest
RUN apk add ca-certificates

COPY --from=build_base /tmp/todo-app-server/out/todo-app-server /app/todo-app-server

EXPOSE 8080

CMD ["/app/todo-app-server"]
