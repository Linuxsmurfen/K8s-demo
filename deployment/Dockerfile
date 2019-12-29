FROM golang:alpine

# Create dir
RUN mkdir /app
WORKDIR /app

# Copy stuff over
COPY demo.go /app/demo.go
COPY demo.html /app/demo.html

# Build webserver
RUN go build -o /app/demo /app/demo.go

# Set enviroment
ENV APP_VERSION="v1.1"
ENV APP_TEXT="Another test.."

# Start webserver
EXPOSE 8080
ENTRYPOINT /app/demo
