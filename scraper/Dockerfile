FROM golang:1.20.1

# Set the Current Working Directory inside the container
WORKDIR /app

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

COPY . .

# Build the application
RUN go build -o main ./src

# Command to run the executable
CMD ["./main"]