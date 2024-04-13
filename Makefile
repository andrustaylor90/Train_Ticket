# Makefile located in the root of TrainTicketServices directory

.PHONY: build run clean docker

build:
	go build -o server ./server/server.go

run:
	go run ./server/server.go

clean:
	rm -f server
	docker stop grpc-train-ticket-server-instance
	docker rm grpc-train-ticket-server-instance
	docker rmi grpc-train-ticket-server

docker:
	./run_docker.sh
