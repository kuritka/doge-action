#!/bin/bash

set -o errexit
set -o pipefail

go mod tidy
go run main.go
