package main

import (
	"log"
	"os"
)

type envOperations struct {
}

func (e *envOperations) exist(key string) (value string) {
	if value = os.Getenv(key); value == "" {
		log.Fatalf("ENV %q is not set.", key)
	}
	return value
}

func (e *envOperations) SetEnv(key string, value string) bool {
	if value == os.Getenv(key) {
		log.Fatalf("Value has already been set : %s", key)
		return false
	}
	os.Setenv(key, value)
	return true
}
