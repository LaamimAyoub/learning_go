package context

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type spyStore struct {
	response string
	t        *testing.T
}

func (s *spyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

// func (s *spyStore) Cancel() {
// 	s.cancelled = true
// }

// func (s *spyStore) assertWasCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Error("store was told to cancel")
// 	}
// }

// func (s *spyStore) assertWasNotCancelled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Error("store was not told to cancel")
// 	}
// }

func TestServer(t *testing.T) {

	t.Run("Normal test", func(t *testing.T) {
		data := "Hello World"
		store := &spyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		// store.assertWasNotCancelled()

	})

	// t.Run("Test Cancel", func(t *testing.T) {
	// 	data := "Hello, world"
	// 	store := &spyStore{response: data, t: t}
	// 	svr := Server(store)

	// 	request := httptest.NewRequest(http.MethodGet, "/", nil)

	// 	cancellingCtx, cancel := context.WithCancel(request.Context())
	// 	time.AfterFunc(5*time.Millisecond, cancel)
	// 	request = request.WithContext(cancellingCtx)

	// 	response := httptest.NewRecorder()

	// 	svr.ServeHTTP(response, request)

	// 	store.assertWasCancelled()

	// })

}
