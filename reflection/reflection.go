package reflection

func Walk(x interface{}, f func(input string)) {
	f("The string should be passing")
}
