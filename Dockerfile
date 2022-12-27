FROM golang:1.18-alpine AS builder

WORKDIR /app

# Copy (go.mod, go.sum) and install dependencies in /app
COPY go.mod go.sum ./
RUN go mod download

# copy sourse code to /app
COPY . .

WORKDIR /app

# Build app
RUN go build -o main .

# ---------- multi-stage --------------
FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/main /app/main

EXPOSE 8081

# Run the binary program produced by `go install`
CMD ["/app/main"]