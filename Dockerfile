# build base go image
FROM golang:1.22.1-alpine as builder

RUN mkdir "/app"

COPY . /app

WORKDIR /app

RUN go mod tidy
RUN go mod download

RUN CGO_ENABLE=0 go build -o shopAPI ./cmd/api/

RUN chmod +x /app/shopAPI

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/shopAPI /app/
COPY .env /app/

CMD [ "/app/shopAPI" ]
