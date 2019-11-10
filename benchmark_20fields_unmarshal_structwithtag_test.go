package json_benchmark

import (
	"testing"
	"encoding/json"

	"github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
)

func Test_20Fields_Unmarshal_StructWithTag(t *testing.T) {
	f := func(t *testing.T, cb func(data []byte, o interface{}) error) {
		var o twentyFieldsStructWithTag
		err := cb(twentyFieldsByte, &o)
		require.NoError(t, err)
		require.Equal(t, twentyFieldsStructWithTagResult, o)
	}
	t.Run(pkgJson, func(t *testing.T) {
		f(t, func(data []byte, o interface{}) error {
			return json.Unmarshal(data, o)
		})
	})
	t.Run(pkgJsoniter, func(t *testing.T) {
		f(t, func(data []byte, o interface{}) error {
			return jsoniter.Unmarshal(data, o)
		})
	})
	t.Run(pkgJsoniterCompatible, func(t *testing.T) {
		f(t, func(data []byte, o interface{}) error {
			return jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, o)
		})
	})
}

func Benchmark_20Fields_Unmarshal_StructWithTag(b *testing.B) {
	b.Run(pkgJson, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o twentyFieldsStructWithTag
			json.Unmarshal(twentyFieldsByte, &o)
		}
	})
	b.Run(pkgJsoniter, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o twentyFieldsStructWithTag
			jsoniter.Unmarshal(twentyFieldsByte, &o)
		}
	})
	b.Run(pkgJsoniterCompatible, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o twentyFieldsStructWithTag
			jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(twentyFieldsByte, &o)
		}
	})
}
