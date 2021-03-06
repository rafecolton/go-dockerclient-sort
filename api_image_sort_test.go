package dockersort_test

import (
	"encoding/json"
	"github.com/fsouza/go-dockerclient"
	"net/http"
	"sort"
	"testing"

	. "github.com/rafecolton/go-dockerclient-sort"
)

var body = `[
     {
             "Repository":"base",
             "Tag":"ubuntu-12.10",
             "Id":"b750fe79269d",
             "Created":1364102657
     },
     {
             "Repository":"base",
             "Tag":"ubuntu-quantal",
             "Id":"b750fe79269d",
             "Created":1364102658
     },
     {
             "RepoTag": [
             "ubuntu:12.04",
             "ubuntu:precise",
             "ubuntu:latest"
             ],
             "Id": "8dbd9e392a964c",
             "Created": 1365714795,
             "Size": 131506275,
             "VirtualSize": 131506275
      },
      {
             "RepoTag": [
             "ubuntu:12.10",
             "ubuntu:quantal"
             ],
             "ParentId": "27cf784147099545",
             "Id": "b750fe79269d2e",
             "Created": 1364102658,
             "Size": 24653,
             "VirtualSize": 180116135
      }
]`

func TestSortByCreatedDescending(t *testing.T) {
	var expected []docker.APIImages
	err := json.Unmarshal([]byte(body), &expected)
	if err != nil {
		t.Fatal(err)
	}
	client := newTestClient(&FakeRoundTripper{message: body, status: http.StatusOK})
	images, err := client.ListImages(false)
	if err != nil {
		t.Fatal(err)
	}
	sort.Sort(ByCreatedDescending(images))
	expectedID := "8dbd9e392a964c"
	if images[0].ID != expectedID {
		t.Errorf(
			"wrong image is first when sorting by created descending: expected %q, got %q",
			expectedID,
			images[0].ID,
		)

	}
}

func TestSortByCreatedAscending(t *testing.T) {
	var expected []docker.APIImages
	err := json.Unmarshal([]byte(body), &expected)
	if err != nil {
		t.Fatal(err)
	}
	client := newTestClient(&FakeRoundTripper{message: body, status: http.StatusOK})
	images, err := client.ListImages(false)
	if err != nil {
		t.Fatal(err)
	}
	sort.Sort(ByCreatedAscending(images))
	expectedID := "b750fe79269d"
	if images[0].ID != expectedID {
		t.Errorf(
			"wrong image is first when sorting by created ascending: expected %q, got %q",
			expectedID,
			images[0].ID,
		)

	}
}
