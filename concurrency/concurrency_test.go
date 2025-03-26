package concurrency

import (
	"testing"
	"time"
)

func slowStubWebsireChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkTestWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = ""
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		CheckWebsites(slowStubWebsireChecker, urls)
	}

}
