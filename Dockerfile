# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Andres Baron af.baron10@uniandes.edu.co"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy Go Modules dependency requirements file
COPY go.mod .

# Copy Go Modules expected hashes file
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy all the app sources (recursively copies files and directories from the host into the image)
COPY . .

# Set http port
ENV PORT 8000

EXPOSE 8080
# Build the app
RUN go build

# Run the app
CMD ["./Posts-API-Golang"]
