build:
	protoc -I. --micro_out=.  --go_out=. proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	sudo docker build -t shippy-service-consignment .
run:
	sudo docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 shippy-service-consignment