#!/bin/bash

echo "physicalcpu_max: $(sysctl -n hw.physicalcpu_max)"
echo "logicalcpu_max : $(sysctl -n hw.logicalcpu_max)"
echo

echo GOMAXPROCS=1 go run ./surface.go
time GOMAXPROCS=1 go run ./surface.go > /dev/null
echo

echo GOMAXPROCS=2 go run ./surface.go
time GOMAXPROCS=2 go run ./surface.go > /dev/null
echo

echo GOMAXPROCS=4 go run ./surface.go
time GOMAXPROCS=4 go run ./surface.go > /dev/null
echo

echo GOMAXPROCS=8 go run ./surface.go
time GOMAXPROCS=8 go run ./surface.go > /dev/null
echo

# echo GOMAXPROCS=1 go test -bench . -benchmem
# GOMAXPROCS=1 go test -bench . -benchmem
# echo

# echo GOMAXPROCS=2 go test -bench . -benchmem
# GOMAXPROCS=2 go test -bench . -benchmem
# echo

# echo GOMAXPROCS=4 go test -bench . -benchmem
# GOMAXPROCS=4 go test -bench . -benchmem
# echo

# echo GOMAXPROCS=8 go test -bench . -benchmem
# GOMAXPROCS=8 go test -bench . -benchmem
# echo