FROM golang:1.21
WORKDIR /src
COPY cmd/gophermart/main.go cmd/gophermart/
COPY internal internal/
COPY go.mod .
COPY go.sum .
RUN go mod download
ENTRYPOINT go test github.com/oktavarium/gomart/internal/app/tests -count=1 -v
