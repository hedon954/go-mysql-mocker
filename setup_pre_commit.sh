#!/bin/bash

set -e

# pre-commit
if ! command -v pre-commit &> /dev/null
then
    echo "pre-commit not found, installing..."
    brew install pre-commit
else
    echo "pre-commit is already installed"
fi


# golangci-lint
if ! command -v golangci-lint &> /dev/null
then
    echo "Installing golangci-lint..."
     go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1
else
    echo "golangci-lint is already installed"
fi


# typos
if ! command -v typos &> /dev/null
then
    echo "Installing typos..."
    brew install typos-cli
else
    echo "typos is already installed"
fi

# pre-commit hooks
echo "Installing pre-commit hooks..."
pre-commit install

echo "All done! pre-commit hooks are installed and configured."
