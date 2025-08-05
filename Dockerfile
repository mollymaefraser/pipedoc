
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o pipedoc ./cmd/pipedoc/


FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/pipedoc /usr/local/bin/pipedoc
ENTRYPOINT ["pipedoc"]
