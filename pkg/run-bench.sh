#!/bin/sh

cd tests/benchmark

go test -bench=. -benchtime=20s -benchmem