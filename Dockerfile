FROM golang:1.14-alpine AS builder

RUN apk update && \
  apk upgrade && \
  apk add --no-cache build-base bash git

WORKDIR /home/app

COPY ./api .

RUN go mod download

RUN go build -o main .

FROM alpine:3.10

WORKDIR /home/app
COPY --from=builder /home/app .
EXPOSE 8080
ENTRYPOINT ["./main"]