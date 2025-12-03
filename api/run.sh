#!/usr/bin/env bash

docker compose up -d 
go run ./cmd/api/v1/main.go
