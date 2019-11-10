package json_benchmark

import (
	"testing"

	"github.com/buger/jsonparser"
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
)

func Test_10Fields_Iterator_Object(t *testing.T) {
	t.Run(pkgJsoniterReadObj, func(t *testing.T) {
		var o tenFieldsStructWithTag
		jit := jsoniter.ConfigDefault.BorrowIterator(tenFieldsByte)
		for field := jit.ReadObject(); field != ""; field = jit.ReadObject() {
			switch field {
			case "intf":
				o.Intf = jit.ReadInt()
			case "uintf":
				o.Uintf = jit.ReadUint()
			case "stringf":
				o.Stringf = jit.ReadString()
			case "uuid":
				o.Uuid = jit.ReadString()
			case "ip":
				o.Ip = jit.ReadString()
			case "email":
				o.Email = jit.ReadString()
			case "int32f":
				o.Int32f = jit.ReadInt32()
			case "uint32f":
				o.Uint32f = jit.ReadUint32()
			case "int64f":
				o.Int64f = jit.ReadInt64()
			case "uint64f":
				o.Uint64f = jit.ReadUint64()
			default:
				require.FailNow(t, "unknown field %q", field)
			}
		}
		require.NoError(t, jit.Error)
		require.Equal(t, tenFieldsStructWithTagResult, o)
	})
	t.Run(pkgJsoniterReadObjCB, func(t *testing.T) {
		var o tenFieldsStructWithTag
		jit := jsoniter.ConfigDefault.BorrowIterator(tenFieldsByte)
		jit.ReadObjectCB(func(it *jsoniter.Iterator, field string) bool {
			switch field {
			case "intf":
				o.Intf = jit.ReadInt()
			case "uintf":
				o.Uintf = jit.ReadUint()
			case "stringf":
				o.Stringf = jit.ReadString()
			case "uuid":
				o.Uuid = jit.ReadString()
			case "ip":
				o.Ip = jit.ReadString()
			case "email":
				o.Email = jit.ReadString()
			case "int32f":
				o.Int32f = jit.ReadInt32()
			case "uint32f":
				o.Uint32f = jit.ReadUint32()
			case "int64f":
				o.Int64f = jit.ReadInt64()
			case "uint64f":
				o.Uint64f = jit.ReadUint64()
			default:
				require.FailNow(t, "unknown field %q", field)
			}
			return true
		})
		require.NoError(t, jit.Error)
		require.Equal(t, tenFieldsStructWithTagResult, o)
	})
	t.Run(pkgJsonparser, func(t *testing.T) {
		var o tenFieldsStructWithTag
		err := jsonparser.ObjectEach(tenFieldsByte, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			var (
				v   int64
				err error
			)
			switch string(key) {
			case `intf`:
				v, err = jsonparser.ParseInt(value)
				o.Intf = int(v)
			case `uintf`:
				v, err = jsonparser.ParseInt(value)
				o.Uintf = uint(v)
			case `stringf`:
				o.Stringf, err = jsonparser.ParseString(value)
			case `uuid`:
				o.Uuid, err = jsonparser.ParseString(value)
			case `ip`:
				o.Ip, err = jsonparser.ParseString(value)
			case `email`:
				o.Email, err = jsonparser.ParseString(value)
			case `int32f`:
				v, err = jsonparser.ParseInt(value)
				o.Int32f = int32(v)
			case `uint32f`:
				v, err = jsonparser.ParseInt(value)
				o.Uint32f = uint32(v)
			case `int64f`:
				v, err = jsonparser.ParseInt(value)
				o.Int64f = int64(v)
			case `uint64f`:
				v, err = jsonparser.ParseInt(value)
				o.Uint64f = uint64(v)
			default:
				require.FailNow(t, "unknown field %q", key)
			}
			return err
		})
		require.NoError(t, err)
		require.Equal(t, tenFieldsStructWithTagResult, o)
	})
}

func Benchmark_10Fields_Iterator_Object(b *testing.B) {
	jit := jsoniter.ConfigDefault.BorrowIterator(nil)
	b.Run(pkgJsoniterReadObj, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o tenFieldsStructWithTag
			jit.ResetBytes(tenFieldsByte)
			for field := jit.ReadObject(); field != ""; field = jit.ReadObject() {
				switch field {
				case "intf":
					o.Intf = jit.ReadInt()
				case "uintf":
					o.Uintf = jit.ReadUint()
				case "stringf":
					o.Stringf = jit.ReadString()
				case "uuid":
					o.Uuid = jit.ReadString()
				case "ip":
					o.Ip = jit.ReadString()
				case "email":
					o.Email = jit.ReadString()
				case "int32f":
					o.Int32f = jit.ReadInt32()
				case "uint32f":
					o.Uint32f = jit.ReadUint32()
				case "int64f":
					o.Int64f = jit.ReadInt64()
				case "uint64f":
					o.Uint64f = jit.ReadUint64()
				}
			}
		}
	})
	b.Run(pkgJsoniterReadObjCB, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o tenFieldsStructWithTag
			jit.ResetBytes(tenFieldsByte)
			jit.ReadObjectCB(func(it *jsoniter.Iterator, field string) bool {
				switch field {
				case "intf":
					o.Intf = jit.ReadInt()
				case "uintf":
					o.Uintf = jit.ReadUint()
				case "stringf":
					o.Stringf = jit.ReadString()
				case "uuid":
					o.Uuid = jit.ReadString()
				case "ip":
					o.Ip = jit.ReadString()
				case "email":
					o.Email = jit.ReadString()
				case "int32f":
					o.Int32f = jit.ReadInt32()
				case "uint32f":
					o.Uint32f = jit.ReadUint32()
				case "int64f":
					o.Int64f = jit.ReadInt64()
				case "uint64f":
					o.Uint64f = jit.ReadUint64()
				}
				return true
			})
		}
	})
	b.Run(pkgJsonparser, func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var o tenFieldsStructWithTag
			jsonparser.ObjectEach(tenFieldsByte, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
				switch string(key) {
				case `intf`:
					v, _ := jsonparser.ParseInt(value)
					o.Intf = int(v)
				case `uintf`:
					v, _ := jsonparser.ParseInt(value)
					o.Uintf = uint(v)
				case `stringf`:
					o.Stringf, _ = jsonparser.ParseString(value)
				case `uuid`:
					o.Uuid, _ = jsonparser.ParseString(value)
				case `ip`:
					o.Ip, _ = jsonparser.ParseString(value)
				case `email`:
					o.Email, _ = jsonparser.ParseString(value)
				case `int32f`:
					v, _ := jsonparser.ParseInt(value)
					o.Int32f = int32(v)
				case `uint32f`:
					v, _ := jsonparser.ParseInt(value)
					o.Uint32f = uint32(v)
				case `int64f`:
					v, _ := jsonparser.ParseInt(value)
					o.Int64f = int64(v)
				case `uint64f`:
					v, _ := jsonparser.ParseInt(value)
					o.Uint64f = uint64(v)
				}
				return nil
			})
		}
	})

}
