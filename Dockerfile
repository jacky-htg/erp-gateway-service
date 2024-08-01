FROM golang:alpine

WORKDIR /app
COPY gateway-service /app/

CMD ["/app/gateway-service"]
