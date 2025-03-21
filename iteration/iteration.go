package iteration

import "strings"

func repeat(c string, n int) string {
	var res strings.Builder
	for i := 0; i < n; i++ {
		res.WriteString(c)
	}
	return res.String()
}

func ExampleRepeat() {
	repeat("a", 4)
	//output: "aaaa"
}
