make:
	go build -o gophermart cmd/gophermart/main.go && ./gophermart

dc:
	docker-compose up

clean:
	rm gophermart
