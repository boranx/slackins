package main

import (
	"gopkg.in/jarcoal/httpmock.v1"
	"testing"
)

func TestTrigger(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://test:8080/job/test/buildWithParameters?token=test&application=ios",
		httpmock.NewStringResponder(201, `[{}]`))

	achieve := trigger("http://test:8080/job/test/buildWithParameters", "test", "&application=ios")
	if !achieve {
		t.Fatalf("jenkins trigger job failed Expected : true, got : False")
	}
}
