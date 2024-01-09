# --- compile the binaries ---------------
FROM golang:alpine AS builder

# Copy stuff over
RUN mkdir /app
COPY demo.go /app/demo.go

# Compile
RUN go build -o /app/demo /app/demo.go

# --- build the final container ---------------
FROM alpine

# Copy stuff over
RUN mkdir /app
COPY demo.html /app/demo.html
COPY --from=builder /app/demo /app/demo

# Set enviroment
ENV APP_VERSION="v1.0"
ENV APP_TEXT="Testing..."

# Start webserver
WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["/app/demo"]
