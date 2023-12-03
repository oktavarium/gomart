make:
	go build -o gophermart cmd/gophermart/main.go && ./gophermart

e2e:
	TEST_MODE=true docker-compose up
	go test github.com/oktavarium/gomart/internal/app/tests -count=1

clean:
	rm gophermart
