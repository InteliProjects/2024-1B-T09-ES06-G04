# Stage 1: Build stage
FROM golang:latest as build

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the entire application and build it
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/avatars ./cmd/main.go

# Stage 2: Final stage
FROM alpine:latest

# Copy the built executable from the build stage
COPY --from=build /app/avatars /app/avatars

# Copy the .env file to the container
COPY .env /app/.env

# Set the working directory
WORKDIR /app

# Expose the port your application runs on
EXPOSE 8081

# Set the command to run your binary executable
CMD ["./avatars"]

