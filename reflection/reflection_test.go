package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	expected := "Chriss"

	var got []string

	x := struct {
		name string
	}{expected}

	Walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("Wrong number of function calls, got %d, want %d", len(got), 1)

	}

	if got[0] != expected {
		t.Errorf("Wrong func call want %s, got %s", expected, got[0])
	}

}

func TestWalkV2(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		//This is case 1
		{
			"struct with one field",
			struct {
				Name string
			}{"Chriss"},
			[]string{"Chriss"},
		},

		//This is case 2
		{
			"struct with two Fields",
			struct {
				Name   string
				Batata string
			}{"Chriss", "Maticha"},
			[]string{"Chriss", "Maticha"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("Failed test")
			}

		})
	}

}
