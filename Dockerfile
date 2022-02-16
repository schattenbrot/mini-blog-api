#
# Build go app
#
FROM golang:1.17 AS builder

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./

# 
# Run app
# 
FROM scratch

WORKDIR /run
COPY .env .
COPY --from=builder /go/src/api ./
ENTRYPOINT [ "./api" ]
