package json_benchmark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_10Fields_Unmarshal_Interface(t *testing.T) {
	Test_Unmarshal_Interface(t, func(t *testing.T, cb func(data []byte) (o interface{}, err error)) {
		o, err := cb(tenFieldsByte)
		require.NoError(t, err)
		m, ok := o.(map[string]interface{})
		require.True(t, ok)
		require.Equal(t, tenFieldsByteMap, m)
	})
}

func Benchmark_10Fields_Unmarshal_Interface(b *testing.B) {
	Benchmark_Unmarshal_Interface(b, tenFieldsByte)
}
