
.PHONY: api
api:
	set -a; \
	. ./app.env; \
	cd ./api; \
	go build -o ./bin/api.exe ./cmd/main.go; \
	./bin/api.exe

.PHONY: pinger
pinger:
	source ./app.env
	go build -o ./pinger/bin/api ./pinger/cmd/main.go
	./pinger/bin/api.exe