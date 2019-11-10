package json_benchmark

import (
	"testing"
	"encoding/json"

	"github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
	"github.com/a8m/djson"
)

func Test_10Fields_Unmarshal_Interface(t *testing.T) {
	f := func(t *testing.T, cb func(data []byte) (interface{}, error)) {
		o, err := cb(tenFieldsByte)
		require.NoError(t, err)
		m, ok := o.(map[string]interface{})
		require.True(t, ok)
		require.Equal(t, tenFieldsByteMap, m)
	}
	t.Run(pkgJson, func(t *testing.T) {
		f(t, func(data []byte) (interface{}, error) {
			var o interface{}
			return o, json.Unmarshal(data, &o)
		})
	})
	t.Run(pkgJsoniter, func(t *testing.T) {
		f(t, func(data []byte) (interface{}, error) {
			var o interface{}
			return o, jsoniter.Unmarshal(data, &o)
		})
	})
	t.Run(pkgJsoniterCompatible, func(t *testing.T) {
		f(t, func(data []byte) (interface{}, error) {
			var o interface{}
			return o, jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, &o)
		})
	})
	t.Run(pkgDjson, func(t *testing.T) {
		f(t, func(data []byte) (interface{}, error) {
			return djson.Decode(data)
		})
	})
}

func Benchmark_10Fields_Unmarshal_Interface(b *testing.B) {
	b.Run(pkgJson, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o interface{}
			json.Unmarshal(tenFieldsByte, &o)
		}
	})
	b.Run(pkgJsoniter, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o interface{}
			jsoniter.Unmarshal(tenFieldsByte, &o)
		}
	})
	b.Run(pkgJsoniterCompatible, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o interface{}
			jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(tenFieldsByte, &o)
		}
	})
	b.Run(pkgDjson, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			djson.Decode(tenFieldsByte)
		}
	})
}
