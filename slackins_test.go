package main

import (
	"gopkg.in/jarcoal/httpmock.v1"
	"testing"
)

func TestTrigger(t *testing.T) {
	jenkins := DefaultJenkins()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", jenkins.uriwithparameters(),
		httpmock.NewStringResponder(201, `[{}]`))

	achieve := trigger(*jenkins)
	if !achieve {
		t.Fatalf("jenkins trigger job failed Expected : true, got : False")
	}
}
