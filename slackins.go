package main

import (
	"log"
	"net/http"
)

func trigger(url string, token string, parameters string) bool {
	resp, err := http.Get(url + "?token=" + token + parameters)
	if err == nil && resp.StatusCode == 201 {
		return true
	} else {
		log.Panicf("Error during trigger process : %s", &err)
		return false
	}
}

func main() {
	env := envOperations{}

	env.SetEnv("jenkins-url", "http://jenkins.trendyol.com:8080/job/test/buildWithParameters")
	env.SetEnv("token", "test")
	env.SetEnv("parameters", "&application=android")

	trigger(env.exist("jenkins-url"), env.exist("token"), env.exist("parameters"))
}
