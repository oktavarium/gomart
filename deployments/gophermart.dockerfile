FROM golang:1.21 as build
WORKDIR /src
COPY cmd/gophermart/main.go cmd/gophermart/main.go
COPY internal internal
COPY go.* .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o gophermart cmd/gophermart/main.go

FROM scratch as run
EXPOSE 8080
COPY --from=build /src/gophermart .
ENTRYPOINT ["/gophermart"]

