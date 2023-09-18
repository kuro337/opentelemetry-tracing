FROM golang:1.21-alpine3.18 as builder

WORKDIR /app

# Copy the go module and sum files to download dependencies

COPY src/go.mod src/go.sum /app/
RUN go mod download

# Copy the source files

COPY src/ /app/

# Build the application with optimizations

RUN go build -o /app/main .

# Use a smaller base image for the runtime

FROM golang:1.21-alpine3.18

# Copy the built binary from the builder stage

COPY --from=builder /app/main /app/main

EXPOSE 8080

CMD ["/app/main"]


# Specify the command to run on container start

#CMD ["/app/main"]

# BUILD
# docker build -t go-fib:latest .

# RUN
#  docker run --name gofib -d -p 8080:8080 go-fib:latest

# Enter Container
# docker run --rm -it go-fib:latest /bin/sh



