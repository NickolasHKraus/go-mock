package main

import (
	"net/http"
)

var URL = "https://example.com"

// TODO
var Get = http.Get

/////////////
// Approach 0
/////////////

// Mock the default Client (DefaultClient) of the http package.
//
// DefaultClient is the default Client and is used by Get, Head, and Post. Its
// name is exported, so it can be replace with a mock client that implements
// its own Get, Head, and Post methods.
//
// NOTE: Get, Head, and Post are simply convience function that issue a request
// with their respective verb to the specified URL.

// This method calls the `Get` method.
func Call0() (*http.Response, error) {
	return http.Get(URL)
}

/////////////
// Approach 1
/////////////

// Mock an http.Client by defining an interface with a `Do` method, which is
// implemented by both the http.Client and the mock.
//
// Disadvantages:
//   * Requires additional complexity to support testing.
//   * The HTTP client (`Client`) is exported, which means it can be set
//     anywhere the module is imported.

type HTTPClient interface {
	// NOTE: This approach requires that the common interface specify the exact
	// method signatures on the http.Client struct.
	//
	// See: net/http/client.go
	Do(req *http.Request) (*http.Response, error)
}

// Client is a package-level variable that implements the HTTPClient interface.
//
// NOTE: This approach requires that our `Client` is exported, so that it can
// be referred to in our test function.
//
// Example:
//
//	```
//	Client = &MockClient{...}
//	```
var Client HTTPClient

// The init function sets `Client` to an instance of the http.Client struct.
//
// NOTE: The init function *is called after all the variable declarations in
// the package have evaluated their initializers...*
//
// See: https://go.dev/doc/effective_go#init
func init() {
	Client = &http.Client{}
}

// This method simply calls the `Do` method on the http.Client struct.
func Call1() (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	// Optionally, set request headers.
	req.Header.Add("Accept", "application/json")
	if err != nil {
		return nil, err
	}
	return Client.Do(req)
}

/////////////
// Approach 2
/////////////

// Use the standard library httptest package, which provides an idiomatic
// method of HTTP testing.
//
// See: https://pkg.go.dev/net/http/httptest
//
// Test case examples can be found here:
//   * https://go.dev/src/net/http/httptest/example_test.go
//
// Disadvantages:
//   * The code under test must allow the URL to be configured to that of the
//     test server.

// This method simply calls the `Do` method on the http.Client struct.
func Call2() (*http.Response, error) {
	c := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	// Optionally, set request headers.
	req.Header.Add("Accept", "application/json")
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func main() {}
