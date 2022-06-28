all: run-kafka-server run-grpc-server

run-kafka-server:
	docker-compose -f ./resources/docker-compose.yaml up -d
	docker exec broker kafka-topics --bootstrap-server broker:9092 --create --topic dblogs

run-grpc-server:
	go run server/grpc_server.go

send-log-to-endpoint:
	curl --header "Content-Type: application/json" --request POST  --data '{"DataDomainID":"5", "Database":"accuknox","Operation":"select", "OperationDetails":"none", "Timestamp":"2022-06-17T09:35:38Z","User":"ada_demo_user"}' http://127.0.0.1:8080/log

demo:
	go run main.go

clean:
	docker-compose -f ./resources/docker-compose.yaml down

