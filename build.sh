#!/bin/bash

# Build the project
echo "Building the project..."
go mod download
go build -o vue-project-backend main.go
echo "Build complete!"