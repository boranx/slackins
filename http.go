package main

import (
	"net/http"
)

func trigger(obj Jenkins) bool {
	resp, err := http.Get(obj.uriwithparameters())
	if err == nil && resp.StatusCode == 201 {
		return true
	} else {
		return false
	}
}
