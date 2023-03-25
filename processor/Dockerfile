FROM golang:1.20.1

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Build the application
RUN go build -o main ./src

# Expose port 8080 to the outside world
EXPOSE 8443

ENV SSL_LOCATION="./certs/"

# Command to run the executable
CMD ["./main"]