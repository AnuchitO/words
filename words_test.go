package words

import (
	"reflect"
	"testing"
)

func TestWords(t *testing.T) {
	tests := []struct {
		in  string
		out map[string]int
	}{
		{
			in: "Nong Nong Nong",
			out: map[string]int{
				"Nong": 3,
			},
		},
		{
			in: "    Nong  Nong ",
			out: map[string]int{
				"Nong": 2,
			},
		},
		{
			in: " Nong  Anuchit Nong",
			out: map[string]int{
				"Nong":    2,
				"Anuchit": 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			result := WordCount(tt.in)
			if !reflect.DeepEqual(tt.out, result) {
				t.Errorf("want %+v, got %+v", tt.out, result)
			}
		})
	}
}
