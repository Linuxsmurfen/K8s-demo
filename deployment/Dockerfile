FROM golang:alpine AS builder

# Create dir
RUN mkdir /app
WORKDIR /app

# Copy stuff over
COPY demo.go /app/demo.go

# Build webserver
RUN go build -o /app/demo /app/demo.go

# --- build final container ----------------------
FROM scratch

# Copy our static executable.
COPY --from=builder /app/demo /app/demo
COPY demo.html /app/demo.html

# Set enviroment
ENV APP_VERSION="v1.0"
ENV APP_TEXT="Testing..."

# Start webserver
EXPOSE 8080
ENTRYPOINT ["/app/demo"]
