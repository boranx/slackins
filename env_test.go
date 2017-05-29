package main

import (
	"os"
	"testing"
)

func Testexist(t *testing.T) {
	env := envOperations{}
	expected := "key"
	value := "value"
	os.Setenv(expected, value)
	achieve := env.exist(expected)
	if achieve != value {
		t.Fatalf("Expected output : %s got : %s", value, achieve)
	}
}

func TestSetEnv(t *testing.T) {
	env := envOperations{}
	expected := "key"
	value := "value"
	if !env.SetEnv(expected, value) {
		t.Fatalf("Expected output :true got : false")
	}
}
