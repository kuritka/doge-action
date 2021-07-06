#!/bin/bash

set -o errexit
set -o pipefail

sudo go mod tidy
sudo go run main.go
