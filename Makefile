make:
	go build -o gophermart cmd/gophermart/main.go && ./gophermart

e2e:
	docker-compose rm -f gophermart client
	TEST_MODE=true docker-compose up --build --abort-on-container-exit

clean:
	rm gophermart
