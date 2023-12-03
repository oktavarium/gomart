FROM golang:1.21 as build
WORKDIR /src
COPY cmd/client/main.go .
COPY internal .
COPY go.* .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o client cmd/client/main.go

FROM scratch as run
EXPOSE 8080
COPY --from=build /src/client .
ENTRYPOINT ["/client"]

