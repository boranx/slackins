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
