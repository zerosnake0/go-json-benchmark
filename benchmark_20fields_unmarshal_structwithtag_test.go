package json_benchmark

import (
	"testing"
	"encoding/json"

	"github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
	ugorji "github.com/ugorji/go/codec"
)

func Test_20Fields_Unmarshal_StructWithTag(t *testing.T) {
	Test_Unmarshal_StructWithTag(t, func(t *testing.T, cb func(data []byte, o interface{}) error) {
		var o twentyFieldsStructWithTag
		err := cb(twentyFieldsByte, &o)
		require.NoError(t, err)
		require.Equal(t, twentyFieldsStructWithTagResult, o)
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
}
