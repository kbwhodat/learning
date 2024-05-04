package main

import (
	"testing"
	"testcontainers/containers"
)

func main() {
	t := &testing.T{}
	containers.TestWithRedis(t)
}
