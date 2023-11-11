FROM  alpine
ENV ACCRUAL_ADDR: "localhost:8081"
ENV DB_URI: "postgres://user:test@localhost:5432/market?sslmode=disable"
COPY cmd/accrual/accrual_darwin_arm64 /bin
CMD ["/bin/accrual_darwin_arm64"]
