FROM golang:1.23 AS dependencies
WORKDIR /pinger
COPY /pinger/go.mod .
COPY /pinger/go.sum .
RUN go mod download


FROM dependencies AS build
WORKDIR /pinger
COPY pinger/. .
RUN go build -o bin/pinger cmd/main.go


FROM debian:stable-slim
WORKDIR /pinger
COPY --from=build /pinger/bin/pinger .
COPY wait-for-it.sh .
RUN chmod +x /pinger/pinger
CMD ["/pinger/wait-for-it.sh","api:8086","--","/pinger/pinger"]