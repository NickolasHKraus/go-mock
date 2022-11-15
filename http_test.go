package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/////////////
// Approach 0
/////////////

type MockRoundTripper struct{}

func (rt MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	body := io.NopCloser(strings.NewReader("Hello, World!"))
	return &http.Response{
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
		Body:       body,
	}, nil
}

func Test0(t *testing.T) {
	http.DefaultClient = &http.Client{
		Transport: MockRoundTripper{},
	}
	exp := "Hello, World!"
	resp, err := Call0()
	ret, _ := io.ReadAll(resp.Body)
	if string(ret) != exp {
		t.Errorf("bad response.\nExpected: %s.\nGot: %s", exp, ret)
	}
	fmt.Printf("%s", ret)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	// ...
}

/////////////
// Approach 1
/////////////

// Mock an http.Client by defining an interface with a `Do` method, which is
// implemented by both the http.Client and the mock.

// MockClient implements the HTTPClient interface.
//
// The MockClient is used in place of the package-level Client, which uses the
// http.Client struct.
//
// It order to allow the return value of the `Do` method to be configurable,
type MockClient struct {
	// MockDoFunc is a field on the MockClient struct that holds the function to
	// be called by the `Do` method.
	MockDoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDoFunc(req)
}

// NOTE: MockClient implements the HTTPClient interface.
func Test1(t *testing.T) {
	Client = &MockClient{
		MockDoFunc: func(req *http.Request) (*http.Response, error) {
			if req.Method != "GET" {
				t.Errorf("Expected request method 'GET', got: %s", req.Method)
			}
			if req.Header.Get("Accept") != "application/json" {
				t.Errorf("Expected request header 'Accept: application/json header', got: %s", req.Header.Get("Accept"))
			}
			body := ioutil.NopCloser(bytes.NewReader([]byte(`{"key":"value"}`)))
			return &http.Response{
				StatusCode: 200,
				Body:       body,
			}, nil
		},
	}
	resp, err := Call1()
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	// ...
}

/////////////
// Approach 2
/////////////

func Test2(t *testing.T) {
	// A Server is an HTTP server listening on a system-chosen port on the local
	// loopback interface, for use in end-to-end HTTP tests.
	//
	// The URL of the HTTP server is of the form http://ipaddr:port with no
	// trailing slash.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			t.Errorf("Expected request method 'GET', got: %s", req.Method)
		}
		if req.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected request header 'Accept: application/json header', got: %s", req.Header.Get("Accept"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"key":"value"}`))
	}))
	defer ts.Close()
	// It should be noted that HTTP requests must be made to the URL of the HTTP
	// server.
	URL = ts.URL
	resp, err := Call2()
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	// ...
}
