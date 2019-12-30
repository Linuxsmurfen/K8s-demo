# --- compile the binaries ---------------
FROM golang:alpine AS builder

# Copy stuff over
RUN mkdir /app
COPY demo.go /app/demo.go

# Build webserver
RUN go build -o /app/demo /app/demo.go

# --- build the final container ---------------
FROM alpine

# Copy stuff over
COPY --from=builder /app/demo /demo
COPY demo.html demo.html

# Set enviroment
ENV APP_VERSION="v1.0"
ENV APP_TEXT="Testing..."

# Start webserver
EXPOSE 8080
#CMD ["/demo"]
ENTRYPOINT ["demo"]
