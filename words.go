package words

import (
	"strings"
)

func WordCount(s string) map[string]int {
	r := map[string]int{}
	for _, w := range strings.Fields(s) {
		r[w]++
	}
	return r
}
