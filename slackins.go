package main

import (
	"log"
	"net/http"
)

func trigger(obj Jenkins) bool {
	resp, err := http.Get(obj.uriwithparameters())
	if err == nil && resp.StatusCode == 201 {
		return true
	} else {
		log.Panicf("Error during trigger process : %s", &err)
		return false
	}
}

func main() {
	env := envOperations{}

	env.SetEnv("URI", "http://jenkins.trendyol.com:8080")
	env.SetEnv("TOKEN", "test")
	env.SetEnv("JOB", "test")

	jenkins := Construct(env.exist("URI"), env.exist("JOB"), env.exist("TOKEN"), map[string]string{"application": "test", "version": "1.1.1"})

	achieve := trigger(*jenkins)
	if !achieve {
		log.Fatalf("jenkins trigger job failed Expected : true, got : False")
	}
}
