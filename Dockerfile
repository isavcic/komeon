FROM golang:1.12 as builder

LABEL maintainer="Ivan Savcic <isavcic@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/isavcic/komeon

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v ./...

# Install the package
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o /komeon .

FROM scratch

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /komeon /

# Run the executable
CMD ["./komeon"]
