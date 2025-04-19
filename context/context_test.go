package context

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type spyStore struct {
	response  string
	cancelled bool
}

func (s *spyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *spyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	data := "Hello World"
	svr := Server(&spyStore{response: data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	}
}
