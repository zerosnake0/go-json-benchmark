package json_benchmark

import (
	"testing"
	"encoding/json"

	"github.com/json-iterator/go"
	"github.com/a8m/djson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mreiferson/go-ujson"
)

func Test_Unmarshal_interface(t *testing.T, f func(t *testing.T, cb func(data []byte) (o interface{}, err error))) {
	t.Run(pkgJson, func(t *testing.T) {
		f(t, func(data []byte) (o interface{}, err error) {
			err = json.Unmarshal(data, &o)
			return
		})
	})
	t.Run(pkgJsoniter, func(t *testing.T) {
		f(t, func(data []byte) (o interface{}, err error) {
			err = jsoniter.Unmarshal(data, &o)
			return
		})
	})
	t.Run(pkgJsoniterCompatible, func(t *testing.T) {
		f(t, func(data []byte) (o interface{}, err error) {
			err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, &o)
			return
		})
	})
	t.Run(pkgDjson, func(t *testing.T) {
		f(t, func(data []byte) (o interface{}, err error) {
			return djson.Decode(data)
		})
	})
	t.Run(pkgEasyJson, func(t *testing.T) {
		f(t, func(data []byte) (o interface{}, err error) {
			l := jlexer.Lexer{
				Data: data,
			}
			o = l.Interface()
			err = l.Error()
			return
		})
	})
	t.Run(pkgUjson, func(t *testing.T) {
		f(t, func(data []byte) (o interface{}, err error) {
			j := ujson.NewDecoder(defaultUjsonObjectStore, data)
			return j.Decode()
		})
	})
}

func Benchmark_Unmarshal_Interface(b *testing.B, data []byte) {
	b.Run(pkgJson, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o interface{}
			json.Unmarshal(data, &o)
		}
	})
	b.Run(pkgJsoniter, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o interface{}
			jsoniter.Unmarshal(data, &o)
		}
	})
	b.Run(pkgJsoniterCompatible, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o interface{}
			jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, &o)
		}
	})
	b.Run(pkgDjson, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			djson.Decode(data)
		}
	})
	b.Run(pkgEasyJson, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			l := jlexer.Lexer{
				Data: data,
			}
			l.Interface()
		}
	})
	b.Run(pkgUjson, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ujson.NewFromBytes(data)
		}
	})
}
