FROM golang:1.23 AS dependencies
WORKDIR /api
COPY /api/go.mod .
COPY /api/go.sum .
RUN go mod download


FROM dependencies AS build
WORKDIR /api
COPY api/. .
RUN go build -o bin/api cmd/main.go


FROM debian:stable-slim
WORKDIR /api
COPY --from=build /api/bin/api .
COPY wait-for-it.sh .
RUN chmod +x /api/api
CMD ["/api/wait-for-it.sh","dbPostgresql:5432","--","/api/api"]