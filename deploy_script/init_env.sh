#!/bin/bash

# Set the desired Go version
GO_VERSION="1.16.7"
PROTOC_VERSION="3.15.8"
GO_TAR="go${GO_VERSION}.linux-amd64.tar.gz"
PROTOC_ZIP="protoc-${PROTOC_VERSION}-linux-x86_64.zip"

# Check if Go is already installed
if ! which go > /dev/null 2>&1; then
    echo "Installing Go..."
    curl -O "https://dl.google.com/go/${GO_TAR}"
    sudo tar -C /usr/local -xzf "${GO_TAR}"
    echo "Go installed successfully."
else
    echo "Go is already installed."
fi

# Set up Go environment variables
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN

# Check if protoc is already installed
if ! which protoc > /dev/null 2>&1; then
    echo "Installing protoc..."
    curl -OL "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP}"
    unzip -o "${PROTOC_ZIP}" -d protoc3
    sudo mv protoc3/bin/* /usr/local/bin/
    sudo mv protoc3/include/* /usr/local/include/
    rm -f "${PROTOC_ZIP}"
    rm -rf protoc3
    echo "protoc installed successfully."
else
    echo "protoc is already installed."
fi

# Check if protoc-gen-go is already installed
if ! which protoc-gen-go > /dev/null 2>&1; then
    echo "Installing protoc-gen-go..."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    echo "protoc-gen-go installed successfully."
else
    echo "protoc-gen-go is already installed."
fi

echo "Environment setup complete."