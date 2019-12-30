# --- compile the binaries ---------------
FROM golang:alpine AS builder

# Copy stuff over
COPY demo.go /demo.go

# Build webserver
RUN go build -o /demo /demo.go

# --- build the final container ---------------
FROM scratch

# Copy stuff over
COPY --from=builder /demo /app/demo
COPY demo.html /app/demo.html

# Set enviroment
ENV APP_VERSION="v1.0"
ENV APP_TEXT="Testing..."

# Start webserver
WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["/app/demo"]
