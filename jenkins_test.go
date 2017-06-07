package main

import (
	"testing"
)

func TestDefaultJenkins_eachvariable(t *testing.T) {
	config := DefaultJenkins()
	expected := "test"
	if config.Token != expected {
		t.Errorf("expected token for testing is %s but got %s", expected, config.Token)
	}
	expected = "http://localhost:8080"
	if config.Uri != expected {
		t.Errorf("expected job uri is to be %s but got %s", expected, config.Uri)
	}
	expected = "test"
	if config.Job != expected {
		t.Errorf("expected job name is to be %s but got %s", expected, config.Job)
	}
	expected = "test"
	if config.Parameters["application"] != expected {
		t.Errorf("expected parameter is to be %s but got %s", expected, config.Parameters["application"])
	}
	expected = "1.1.1"
	if config.Parameters["version"] != expected {
		t.Errorf("expected parameter is to be %s but got %s", expected, config.Parameters["version"])
	}
	expected = "http://localhost:8080/job/test/buildWithParameters?token=test"
	if config.uri() != expected {
		t.Errorf("expected parameter is to be %s but got %s", expected, config.uri())
	}
	expected = "http://localhost:8080/job/test/buildWithParameters?token=test&application=test&version=1.1.1"
	if config.uriwithparameters() != expected {
		t.Errorf("expected parameter is to be %s but got %s", expected, config.uriwithparameters())
	}
}

func TestConstruct(t *testing.T) {
	test := Construct("http://jenkins:8080", "test", "test", map[string]string{"application": "test", "version": "1.1.1"})
	expected := "http://jenkins:8080/job/test/buildWithParameters?token=test&application=test&version=1.1.1"
	if test.uriwithparameters() != expected {
		t.Errorf("expected parameter is to be %s but got %s", expected, test.uriwithparameters())
	}
}
