package Select

import "net/http"

func RacerV2(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b

	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch

}
