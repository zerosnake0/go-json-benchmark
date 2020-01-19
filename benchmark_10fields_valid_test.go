package json_benchmark

import (
	"testing"
)

func Test_10Fields_Valid(t *testing.T) {
	Test_Valid(t, tenFieldsByte, true)
}

func Benchmark_10Fields_Valid(b *testing.B) {
	Benchmark_Valid(b, tenFieldsByte)
}
