package json_benchmark

import (
	"testing"
)

func Test_20Fields_Valid(t *testing.T) {
	Test_Valid(t, twentyFieldsByte, true)
}

func Benchmark_20Fields_Valid(b *testing.B) {
	Benchmark_Valid(b, twentyFieldsByte)
}
