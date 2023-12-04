make:
	go build -o gophermart cmd/gophermart/main.go && ./gophermart

e2e:
	docker container rm -f gophermart client marketdb
	TEST_MODE="true" docker-compose up  --build --abort-on-container-exit


clean:
	rm gophermart
