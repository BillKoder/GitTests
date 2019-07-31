# Selete the base image
FROM golang:latest

# Copy in out source file
COPY main.go ./main.go

# Compile the code
RUN go build -o /app/hello ./main.go

# Set the default command to be executed when container is started from this image
CMD [ "/app/hello" ]