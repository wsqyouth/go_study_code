#!/bin/bash

# Check if the correct number of arguments was provided
if [ "$#" -ne 3 ]; then
    echo "Usage: $0 <ip> <username> <appname>"
    exit 1
fi

# Assign the provided arguments to variables
IP=$1
USERNAME=$2
APPNAME=$3
DEPLOY_SCRIPT="deploy.go"

# Build the deploy.go script
echo "Building the deploy script..."
go build -o deploy $DEPLOY_SCRIPT
if [ $? -ne 0 ]; then
    echo "Failed to build the deploy script."
    exit 1
fi
echo "Deploy script built successfully."

# Run the deploy executable to build the app binary and deploy it
echo "Building and deploying the app using deploy script..."
./deploy $IP $USERNAME $APPNAME
if [ $? -ne 0 ]; then
    echo "Failed to build and deploy the app."
    exit 1
fi
echo "App built and deployed successfully."