# autodom

Simple test app with `go-kit` and microservice architecture. Based on [github.com/shijuvar/gokit-examples](https://github.com/shijuvar/gokit-examples).

## Run

### Launch on host

	cd services/advisor/cmd
	go build main.go
	./main -u <DB_USER> -p <DB_PASSWORD> -d <DB_NAME> -a <LISTENING_ADDRESS>
	
Application tries to get DB_USER, DB_PASSWORD, DB_NAME from env vars, if cli args is not provided.

### Launch in container

	docker-compose up



 

	