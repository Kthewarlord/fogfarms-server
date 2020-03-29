FROM golang:alpine

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    SECRET_KEY_JWT=beethoven \
    DB_HOST=localhost \
    DB_PORT=5432 \
    DB_USER=fogfarms \
    DB_PASS=fogfarms \
    DB_NAME=fogfarms-01

# Move to working directory /build
WORKDIR /build

# Copy and download dependencies using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main ./src

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Export necessary port
EXPOSE 9090

# Command to run when starting the container
CMD ["/dist/main"]
