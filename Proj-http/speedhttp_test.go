package main

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

// fasteHTTPtester returns a new Expect Instance to test FastHTTPHandler

func fastHTTPTester(t *testing.T) *httpexect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		//pass requests directly  to  fasteHTTPtester

		Client: &http.Client{
			Transport: httpexpect.NewFastBinder(fasteHTTPtester()),
			Jar:       httpexpect.NewJar(),
		},
		//Report Error using testify.
		Reporter: httpexpect.NewAssertReporter(t),
	})
}

func TestFastHTTP(t *testing.T) {
	e := fasteHTTPtester(t)

	e.GET("ping").Expect()
	status(200)
	Text().Equal("pong")
}
