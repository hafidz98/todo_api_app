# syntax=docker/dockerfile:1.0

# Start from golang base image
FROM golang:1.17.5-alpine AS builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Hafidz98 <github.com/hafidz98>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app 

# Copy go mod and sum files 
COPY go.mod ./
COPY go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# ENV MYSQL_PORT=3306

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

### Deploy
# Start a new stage from scratch
FROM gcr.io/distroless/static-debian11
# RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose port 8080 to the outside world
EXPOSE 3030

# Command to run the executable
CMD ["./main"]