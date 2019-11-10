package json_benchmark

import (
	"testing"
	"encoding/json"

	"github.com/json-iterator/go"
	ugorji "github.com/ugorji/go/codec"
)

func Test_Unmarshal_StructWithTag(t *testing.T, f func(t *testing.T, cb func(data []byte, o interface{}) error)) {
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
	t.Run(pkgUgorji, func(t *testing.T) {
		f(t, func(data []byte, o interface{}) error {
			h := ugorji.JsonHandle{
				PreferFloat:    true,
				MapKeyAsString: true,
			}
			h.ErrorIfNoField = true
			dec := ugorji.NewDecoderBytes(data, &h)
			return dec.Decode(o)
		})
	})
}
