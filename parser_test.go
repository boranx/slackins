package main

import "testing"

func TestParser(t *testing.T) {
	key := "param1=1,param2=2"
	expected := "1"
	mapping := *parse(key)
	if mapping["param1"] != expected {
		t.Errorf("expected parameter is to be %s but got %s", expected, mapping["param1"])
	}
	expected = "2"
	if mapping["param2"] != expected {
		t.Errorf("expected parameter is to be %s but got %s", expected, mapping["param2"])
	}
}
