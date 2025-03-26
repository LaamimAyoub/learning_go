package Select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(time.Duration(20))
	fastServer := makeDelayedServer(time.Duration(0))

	fastUrl := fastServer.URL
	slowUrl := slowServer.URL
	want := fastUrl
	got := RacerV2(slowUrl, fastUrl)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}

func makeDelayedServer(secondes time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(secondes * time.Microsecond)
		w.WriteHeader(http.StatusOK)
	}))

}
