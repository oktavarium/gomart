FROM golang:1.21 as build
WORKDIR /src
COPY cmd/gophermart/main.go cmd/gophermart/main.go
COPY internal internal
COPY go.* .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o gophermart cmd/gophermart/main.go

FROM scratch as run
ENV ACCRUAL_ADDR: "localhost:8081"
ENV DB_URI: "postgres://user:test@localhost:5432/market?sslmode=disable"
ENV GOPHERMART_ADDR: "localhost:8080"
COPY --from=build /src/gophermart /run/gophermart
CMD ["/run/gophermart", "-a", ACCRUAL_ADDR, "-d", DB_URI, "-r", GOPHERMART_ADDR]
