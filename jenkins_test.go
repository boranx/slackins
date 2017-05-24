package main

import (
	"testing"
)

func TestDefaultJenkins_token(t *testing.T) {
	config := DefaultJenkins()
	expected := "test"
	if config.Token != expected {
		t.Errorf("expected token for testing is %s but got %s", expected, config.Token)
	}
}

func TestDefaultJenkins_uri(t *testing.T) {
	config := DefaultJenkins()
	expected := "http://localhost:8080"
	if config.Uri != expected {
		t.Errorf("expected job uri is to be %d but got %d", expected, config.Uri)
	}
}

func TestDefaultJenkins_job(t *testing.T) {
	config := DefaultJenkins()
	expected := "test"
	if config.Job != expected {
		t.Errorf("expected job name is to be %d but got %d", expected, config.Job)
	}
}

func TestDefaultJenkins_parameterone(t *testing.T) {
	config := DefaultJenkins()
	expected := "test"
	if config.Parameters["application"] != expected {
		t.Errorf("expected parameter is to be %d but got %d", expected, config.Parameters["application"])
	}
}

func TestDefaultJenkins_parametersecond(t *testing.T) {
	config := DefaultJenkins()
	expected := "1.1.1"
	if config.Parameters["version"] != expected {
		t.Errorf("expected parameter is to be %d but got %d", expected, config.Parameters["version"])
	}
}

func TestUri(t *testing.T) {
	config := DefaultJenkins()
	expected := "http://localhost:8080/job/test/buildWithParameters?token=test"
	if config.uri() != expected {
		t.Errorf("expected parameter is to be %d but got %d", expected, config.uri())
	}
}

func TestUriwithParameters(t *testing.T) {
	config := DefaultJenkins()
	expected := "http://localhost:8080/job/test/buildWithParameters?token=test&application=test&version=1.1.1"
	if config.uriwithparameters() != expected {
		t.Errorf("expected parameter is to be %d but got %d", expected, config.uriwithparameters())
	}
}
