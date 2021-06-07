run:
	@echo "===== RUNNING MUSIC CHART SERVICES ====="
	@go build -o app && ./app

test:
	@go test -v -cover -covermode=atomic ./...

server-up:
	@echo "===== START SERVER ====="
	@sudo docker build -t lp/mysql .
	@sudo docker run --name lp-db -p 3307:3306 -d lp/mysql

server-down:
	@echo "===== STOP SERVER ====="
	@sudo docker container stop lp-db
	@sudo docker container rm lp-db