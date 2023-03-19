build:
	go build -o messenger server/cmd/main.go

start:
	docker-compose up --build go db

test:
	docker-compose up --build -d test_db
	cd ./server/handlers; go test -v; cd ../..
	docker-compose stop test_db
	docker rm -f test_myapp_db
	rm -rf ./test_data/

.PHONY: build start test