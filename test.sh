#!/bin/bash

echo "========== 🔍 Ejecutando tests... ============"
go test ./server -v
go test ./client -v

echo "========== 🏎 Ejecutando benchmarks... ======="
echo ""
go test -bench=. ./server
go test -bench=. ./client
