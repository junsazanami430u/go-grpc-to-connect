#!/bin/bash

go test -cover -v ./...  -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html