package json_benchmark

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
	ugorji "github.com/ugorji/go/codec"
	"github.com/zerosnake0/jzon"
)

func Test_10Fields_Unmarshal_StructWithTag(t *testing.T) {
	Test_Unmarshal_StructWithTag(t, func(t *testing.T, cb func(data []byte, o interface{}) error) {
		var o tenFieldsStructWithTag
		err := cb(tenFieldsByte, &o)
		require.NoError(t, err)
		require.Equal(t, tenFieldsStructWithTagResult, o)
	})
}

func Benchmark_10Fields_Unmarshal_StructWithTag(b *testing.B) {
	b.Run(pkgJson, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o tenFieldsStructWithTag
			json.Unmarshal(tenFieldsByte, &o)
		}
	})
	b.Run(pkgJsoniter, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o tenFieldsStructWithTag
			jsoniter.Unmarshal(tenFieldsByte, &o)
		}
	})
	b.Run(pkgJsoniterCompatible, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o tenFieldsStructWithTag
			jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(tenFieldsByte, &o)
		}
	})
	b.Run(pkgUgorji, func(b *testing.B) {
		b.ReportAllocs()
		h := ugorji.JsonHandle{
			PreferFloat:    true,
			MapKeyAsString: true,
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			dec := ugorji.NewDecoderBytes(tenFieldsByte, &h)
			var o tenFieldsStructWithTag
			dec.Decode(&o)
		}
	})
	b.Run(pkgJzon, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o tenFieldsStructWithTag
			jzon.Unmarshal(tenFieldsByte, &o)
		}
	})
	b.Run(pkgJzonFast, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o tenFieldsStructWithTag
			jzonFastDecoderConfig.Unmarshal(tenFieldsByte, &o)
		}
	})
}
