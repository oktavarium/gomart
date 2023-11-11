FROM golang:1.21 as build
WORKDIR /src
COPY cmd/gophermart/main.go cmd/gophermart/main.go
COPY internal internal
COPY go.* .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o gophermart cmd/gophermart/main.go

FROM scratch as run
EXPOSE 8080
ENV ACCRUAL_SYSTEM_ADDRESS = "localhost:8888"
ENV DATABASE_URI = "postgresql://user:user@postgres:5432/market?sslmode=disable"
ENV RUN_ADDRESS = "0.0.0.0:8080"
COPY --from=build /src/gophermart .
ENTRYPOINT ["/gophermart"]

