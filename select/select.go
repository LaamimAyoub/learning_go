package Select

import (
	"net/http"
	"time"
)

func time_website(a string) time.Duration {
	start := time.Now()
	http.Get(a)
	return time.Since(start)
}

func Racer(a, b string) (winner string) {

	duration_a := time_website(a)
	duration_b := time_website(b)

	if duration_a < duration_b {
		return a
	}
	return b

}
