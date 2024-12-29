#!/bin/bash

# Replace old import paths with new ones
find . -type f -name "*.go" -exec sed -i 's|github.com/your-username/irt-exam-system|irt-exam-system/backend|g' {} +

# Run go mod tidy to update dependencies
go mod tidy 