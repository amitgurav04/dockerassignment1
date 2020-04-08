# Dockerfile References: https://docs.docker.com/engine/reference/builder/
# Start from golang v1.12.0 base image
FROM golang:1.12.0

# Add Maintainer Info
LABEL maintainer="Amit Gurav <amit.gurav@calsoftinc.com>"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/amitgurav04/assignment_client

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v ./... && GO111MODULE=auto go build -o /client .
# Install the package
#RUN go install -v ./...

#RUN GO111MODULE=auto go build .

# This container exposes port 9100 to the outside world
EXPOSE 8090
USER root

# Run the executable
ENTRYPOINT ["/client"]
