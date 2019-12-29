FROM golang:alpine

# Copy stuff over
COPY demo.go /demo.go
COPY demo.html /demo.html

# Build webserver
RUN go build -o demo demo.go

# Set enviroment
ENV APP_VERSION="v1.0"
ENV APP_TEXT="Testing"

# Start webserver
EXPOSE 8080
ENTRYPOINT /demo
