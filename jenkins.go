package main

import "fmt"

type Jenkins struct {
	// Uri for jenkins job
	Uri string

	// Job name that will be triggered
	Job string

	// Token in order to trigger this job
	Token string

	// Job parameters for `buildWithParameters`
	Parameters map[string]string
}

func DefaultJenkins() *Jenkins {
	return &Jenkins{
		Uri:        "http://localhost:8080",
		Job:        "test",
		Token:      "test",
		Parameters: map[string]string{"application": "test", "version": "1.1.1"},
	}
}

func Construct(uri string, job string, token string, parameters map[string]string) *Jenkins {
	jenkins := Jenkins{}
	jenkins.Uri = uri
	jenkins.Job = job
	jenkins.Token = token
	jenkins.Parameters = parameters

	return &jenkins
}

func (jnks *Jenkins) uri() (uri string) {
	return fmt.Sprintf("%s/job/%s/buildWithParameters?token=%s", jnks.Uri, jnks.Job, jnks.Token)
}

func (jnks *Jenkins) uriwithparameters() (uri string) {
	var construct string
	for k := range jnks.Parameters {
		construct = construct + fmt.Sprintf("&%s=%s", k, jnks.Parameters[k])
	}
	return fmt.Sprintf("%s%s", jnks.uri(), construct)
}
