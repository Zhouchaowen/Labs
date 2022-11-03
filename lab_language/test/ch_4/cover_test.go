// https://brantou.github.io/2017/05/24/go-cover-story/
package ch_4

import "testing"

type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "negative"},
	{5, "small"},
}

// go test -v -run=TestSize -cover
// go test -v -run=TestSize -coverprofile=size_coverage.out
func TestSize(t *testing.T) {
	for i, test := range tests {
		size := Size(test.in)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}
