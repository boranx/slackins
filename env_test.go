package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestExist(t *testing.T) {
	env := envOperations{}

	var value string
	key := "TEST"

	if value = os.Getenv(key); value != "" {
		os.Unsetenv(key)
		defer func() { os.Setenv(key, value) }()
	}

	if os.Getenv("TEST_MUST_ENV") == "1" {
		env.exist(key)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestExist")
	cmd.Env = append(os.Environ(), "TEST_MUST_ENV=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Errorf("Exist() failed to fatal if env %q is not present", key)
}

func TestSetEnv(t *testing.T) {
	env := envOperations{}

	var value string
	key := "IDIOT"

	if value = os.Getenv(key); value != "" {
		os.Unsetenv(key)
		defer func() { env.SetEnv(key, value) }()
	}

	if os.Getenv("TEST_MUST_ENV") == "2" {
		env.exist(key)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestSetEnv")
	cmd.Env = append(os.Environ(), "TEST_MUST_ENV=2")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Errorf("Exist() failed to fatal if env %q is not present", key)
}
