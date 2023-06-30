# Use the official Golang image as the base
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Install additional development tools
#RUN go get github.com/githubnemo/CompileDaemon
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

# Copy the source code to the container
COPY . ./

ENTRYPOINT CompileDaemon -log-prefix=false  -build="go build -buildvcs=false -o ./src/my-marketplace ./src" -command="./src/my-marketplace"
