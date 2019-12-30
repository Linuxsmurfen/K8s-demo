# --- compile the binaries ---------------
FROM golang:alpine AS builder

# Copy stuff over
RUN mkdir /app
COPY demo.go /app/demo.go

# Build webserver
WORKDIR /app
RUN go build -o /app/demo /app/demo.go

# --- build the final container ---------------
FROM scratch

# Copy stuff over
COPY --from=builder /app/demo /app/demo
COPY demo.html /app/demo.html

# Set enviroment
ENV APP_VERSION="v1.0"
ENV APP_TEXT="Testing..."

# Start webserver
EXPOSE 8080
ENTRYPOINT ["/app/demo"]
