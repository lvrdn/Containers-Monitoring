compose_up:
	docker-compose --env-file ./app.env up

on:
	docker-compose --env-file ./app.env start

off:
	docker-compose --env-file ./app.env stop

.PHONY: api
api:
	set -a; \
	. ./app.env; \
	cd ./api; \
	go build -o ./bin/api.exe ./cmd/main.go; \
	./bin/api.exe

.PHONY: pinger
pinger:
	set -a; \
	. ./app.env; \
	cd ./pinger; \
	go build -o ./bin/pinger.exe ./cmd/main.go; \
	./bin/pinger.exe

.PHONY: frontend
frontend:
	set -a; \
	. ./app.env; \
	cd ./frontend/react-intro; \
	npm start

db:
	docker run --name dbmonitoring -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -e POSTGRES_DB=monitoring -d postgres 

db_connect:
	psql -hlocalhost -p5432 -Uroot -dmonitoring 