# Stage 1: Build the User Management Service
FROM golang:1.23 as builder
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application source code
COPY . ./

# Build the application binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bank-Services .

# Stage 2: Create a lightweight container
FROM alpine:latest
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/bank-Services .

# Make the binary executable
RUN chmod +x bank-Services

# Set the default command to run the binary
CMD ["./bank-Services"]
