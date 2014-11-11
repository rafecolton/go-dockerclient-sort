package dockersort_test

import (
	"github.com/fsouza/go-dockerclient"
	"io/ioutil"
	"net/http"
	"strings"
)

func newTestClient(rt *FakeRoundTripper) *docker.Client {
	endpoint := "http://localhost:4243"
	client, _ := docker.NewClient(endpoint)
	client.HTTPClient = &http.Client{Transport: rt}
	return client
}

type FakeRoundTripper struct {
	message  string
	status   int
	header   map[string]string
	requests []*http.Request
}

func (rt *FakeRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	body := strings.NewReader(rt.message)
	rt.requests = append(rt.requests, r)
	res := &http.Response{
		StatusCode: rt.status,
		Body:       ioutil.NopCloser(body),
		Header:     make(http.Header),
	}
	for k, v := range rt.header {
		res.Header.Set(k, v)
	}
	return res, nil
}

func (rt *FakeRoundTripper) Reset() {
	rt.requests = nil
}
