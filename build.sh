#!/bin/bash

NAME="eve-trader"

echo "Building..."
go build -o $NAME ./cmd/
echo "Completed."


