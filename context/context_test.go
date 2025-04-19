package context

import (
	"context"
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

	t.Run("Normal test", func(t *testing.T) {
		data := "Hello World"
		store := &spyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		if store.cancelled {
			t.Error("store should not be canceld")
		}

	})

	t.Run("Test Cancel", func(t *testing.T) {
		data := "Hello, world"
		store := &spyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if !store.cancelled {
			t.Error("store was not told to cancel")
		}

	})

}
