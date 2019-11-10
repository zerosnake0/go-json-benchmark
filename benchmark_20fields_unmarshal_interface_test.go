package json_benchmark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_20Fields_Unmarshal_Interface(t *testing.T) {
	Test_Unmarshal_Interface(t, func(t *testing.T, cb func(data []byte) (o interface{}, err error)) {
		o, err := cb(twentyFieldsByte)
		require.NoError(t, err)
		m, ok := o.(map[string]interface{})
		require.True(t, ok)
		require.Equal(t, twentyFieldsByteMap, m)
	})
}

func Benchmark_20Fields_Unmarshal_Interface(b *testing.B) {
	Benchmark_Unmarshal_Interface(b, twentyFieldsByte)
}
