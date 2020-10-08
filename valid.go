package json_benchmark

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"github.com/zerosnake0/jzon"
)

func Test_Valid(t *testing.T, data []byte, valid bool) {
	t.Run(pkgJson, func(t *testing.T) {
		require.Equal(t, valid, json.Valid(data))
	})
	t.Run(pkgJsoniter, func(t *testing.T) {
		require.Equal(t, valid, jsoniter.Valid(data))
	})
	t.Run(pkgJsoniterCompatible, func(t *testing.T) {
		require.Equal(t, valid, jsoniter.ConfigCompatibleWithStandardLibrary.Valid(data))
	})
	t.Run(pkgGJson, func(t *testing.T) {
		require.Equal(t, valid, gjson.ValidBytes(data))
	})
	t.Run(pkgJzon, func(t *testing.T) {
		require.Equal(t, valid, jzon.Valid(data))
	})
}

func Benchmark_Valid(b *testing.B, data []byte) {
	b.Run(pkgJson, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			json.Valid(data)
		}
	})
	b.Run(pkgJsoniter, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			jsoniter.Valid(data)
		}
	})
	b.Run(pkgJsoniterCompatible, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			jsoniter.ConfigCompatibleWithStandardLibrary.Valid(data)
		}
	})
	b.Run(pkgGJson, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			gjson.ValidBytes(data)
		}
	})
	b.Run(pkgJzon, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			jzon.Valid(data)
		}
	})
}
