package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	var theTests = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"404", "/doesnotexist", http.StatusNotFound},
	}
	var app application
	routes := app.routes()

	// create a test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	pathToTemplates = "./../../templates/"
	// range through tests
	for _, test := range theTests {
		resp, err := ts.Client().Get(ts.URL + test.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}
		if resp.StatusCode != test.expectedStatusCode {
			t.Errorf("for %s: expected was %d but got %d", test.name, test.expectedStatusCode, resp.StatusCode)
		}
	}

}
