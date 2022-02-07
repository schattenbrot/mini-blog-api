#
# Build go app
#
FROM golang:1.17 AS builder

WORKDIR /go/src
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o api ./cmd/api

# 
# Run app
# 
FROM alpine:latest

RUN apk add --no-cache libc6-compat

RUN addgroup -S apiUser && adduser -S apiUser -G apiUser
USER apiUser

WORKDIR /run
COPY .env .
COPY --from=builder /go/src/api ./
ENTRYPOINT [ "./api" ]