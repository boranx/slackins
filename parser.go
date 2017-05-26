package main

import "strings"

func parse(key string) *map[string]string {
	params := make(map[string]string)
	result := strings.Split(key, ",")
	for i := range result {
		parsed := strings.Split(result[i], "=")
		params[parsed[0]] = parsed[1]
	}
	return &params
}
