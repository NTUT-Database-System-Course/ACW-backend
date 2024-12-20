FROM golang:1.22-bullseye

# Install necessary packages
RUN apt-get update && apt-get install -y \
    vim \
    git \
    && rm -rf /var/lib/apt/lists/*

# Set environment variables
ENV GO111MODULE=on \
    GOPROXY=https://proxy.golang.org,direct

# Install swag for Swagger documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Set working directory
WORKDIR /workspaces

# Copy Go modules files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Generate Swagger documentation
RUN swag init -g main.go --parseDependency github.com/guregu/null/v5

# Expose port (optional, since it's handled in docker-compose)
EXPOSE 8080

# Command to start the application
CMD ["go", "run", "main.go"]