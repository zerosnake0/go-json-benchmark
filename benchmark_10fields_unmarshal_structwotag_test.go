package json_benchmark

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
	"github.com/zerosnake0/jzon"
)

func Test_Unmarshal_10Fields_StructWoTag(t *testing.T) {
	f := func(t *testing.T, cb func(data []byte, o interface{}) error) {
		var o tenFieldsStructWoTag
		err := cb(tenFieldsByte, &o)
		require.NoError(t, err)
		require.Equal(t, tenFieldsStructWoTagResult, o)
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
	t.Run(pkgJzon, func(t *testing.T) {
		f(t, func(data []byte, o interface{}) error {
			return jzon.Unmarshal(data, o)
		})
	})
}

func Benchmark_10Fields_Unmarshal_StructWoTag(b *testing.B) {
	b.Run(pkgJson, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o tenFieldsStructWoTag
			json.Unmarshal(tenFieldsByte, &o)
		}
	})
	b.Run(pkgJsoniter, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o tenFieldsStructWoTag
			jsoniter.Unmarshal(tenFieldsByte, &o)
		}
	})
	b.Run(pkgJsoniterCompatible, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o tenFieldsStructWoTag
			jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(tenFieldsByte, &o)
		}
	})
	b.Run(pkgJzon, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o tenFieldsStructWoTag
			jzon.Unmarshal(tenFieldsByte, &o)
		}
	})
}
