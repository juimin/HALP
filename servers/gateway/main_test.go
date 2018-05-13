package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {

	cases := []struct {
		name          string
		env           string
		defaultVal    string
		errorMessage  string
		expectedVar   string
		expectedError error
	}{
		{
			name:          "Test Gopath - We already know this is set",
			env:           "GOPATH",
			defaultVal:    "",
			errorMessage:  "Go Path Not Set",
			expectedVar:   os.Getenv("GOPATH"),
			expectedError: nil,
		},
		{
			name:          "Random Env - No env variable set + no default",
			env:           "Choclate",
			defaultVal:    "",
			errorMessage:  "Choclate Not Set",
			expectedVar:   "",
			expectedError: fmt.Errorf("Choclate Not Set"),
		},
		{
			name:          "Random Env - No env variable set + default set",
			env:           "Choclate",
			defaultVal:    "white",
			errorMessage:  "Choclate Not Set",
			expectedVar:   "white",
			expectedError: nil,
		},
	}

	for _, c := range cases {
		val, err := getEnvVariable(c.env, c.defaultVal, c.errorMessage)

		if err == nil {
			if c.expectedVar != val {
				t.Errorf("%s Failed: Expected %s but got %s", c.name, c.expectedVar, val)
			}
		} else {
			if err.Error() != c.expectedError.Error() {
				t.Errorf("%s Failed: Expected %s but got %s", c.name, c.expectedError, err)
			}
		}
	}
}

func TestMain(t *testing.T) {
	_, _, _, _, err := generateContextHandler()
	if err == nil && os.Getenv("ADDR") == "" {
		t.Errorf("Something should have gone wrong because we don't have all the information")
	}
	os.Setenv("ADDR", "localhost:8080")
	_, _, _, _, err = generateContextHandler()
	if err == nil && os.Getenv("TLSKEY") == "" {
		t.Errorf("Something should have gone wrong because we don't have all the information")
	}
	os.Setenv("TLSKEY", os.Getenv("GOPATH")+"/src/github.com/JuiMin/HALP/servers/gateway/tls/privkey.pem")
	_, _, _, _, err = generateContextHandler()
	if err == nil && os.Getenv("TLSCERT") == "" {
		t.Errorf("Something should have gone wrong because we don't have all the information")
	}
	os.Setenv("TLSCERT", os.Getenv("GOPATH")+"/src/github.com/JuiMin/HALP/servers/gateway/tls/fullchain.pem")
	_, _, _, _, err = generateContextHandler()
	if err == nil && os.Getenv("SESSIONKEY") == "" {
		t.Errorf("Something should have gone wrong because we don't have all the information")
	}
	os.Setenv("SESSIONKEY", "spUPraqUgethu4AF?x")
	_, _, _, _, err = generateContextHandler()

	if err != nil {
		t.Errorf("There was a problem generating the cr")
	}
}

func TestGenerateMux(t *testing.T) {
	os.Setenv("ADDR", "localhost:8080")
	os.Setenv("TLSKEY", os.Getenv("GOPATH")+"/src/github.com/JuiMin/HALP/servers/gateway/tls/privkey.pem")
	os.Setenv("TLSCERT", os.Getenv("GOPATH")+"/src/github.com/JuiMin/HALP/servers/gateway/tls/fullchain.pem")
	os.Setenv("SESSIONKEY", "spUPraqUgethu4AF?x")
	cr, port, tlscert, tlskey, err := generateContextHandler()

	if err != nil {
		t.Errorf("There was a problem generating the cr")
	}

	corsHandler := generateMux(cr, tlscert, tlskey, port)

	if corsHandler == nil {
		t.Errorf("CORS handler can't be nil")
	}
}
