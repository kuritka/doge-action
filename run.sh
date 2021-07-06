#!/bin/bash

set -o errexit
set -o pipefail

pwd
ls -la


sudo go mod tidy
#sudo go run ./main.go
