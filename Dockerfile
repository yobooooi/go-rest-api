FROM golang:1.15

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/go-rest-api/

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN env GOOS=linux GOARCH=amd64 go install -v ./...
RUN ls -l
# This container exposes port 8080 to the outside world
EXPOSE 10000

CMD ["/bin/bash", "-c", "go-rest-api"]