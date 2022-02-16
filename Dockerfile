#
# Build go app
#
FROM golang:1.17 AS builder

WORKDIR /go/src
COPY . .

# Buildkit needs to be enabled in order for TARGETARCH variable to exist on build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -o api ./cmd/api

# 
# Run app
# 
FROM scratch

WORKDIR /run
COPY .env .
COPY --from=builder /go/src/api ./
ENTRYPOINT [ "./api" ]
