package json_benchmark

import (
	"encoding/json"
)

var tenFieldsByte = []byte(`{
	"intf":    -123,
	"uintf":   456,
	"stringf": "abc",
	"uuid":    "de305d54-75b4-431b-adb2-eb6b9e546014",
	"ip":      "127.0.0.1",
	"email":   "test@example.com",
	"int32f":  -164281,
	"uint32f": 2938729,
	"int64f":  -8973,
	"uint64f": 89302174
}`)

var tenFieldsByteMap map[string]interface{}

type tenFieldsStructWoTag struct {
	Intf    int
	Uintf   uint
	Stringf string
	Uuid    string
	Ip      string
	Email   string
	Int32f  int32
	Uint32f uint32
	Int64f  int64
	Uint64f uint64
}

var tenFieldsStructWoTagResult tenFieldsStructWoTag

type tenFieldsStructWithTag struct {
	Intf    int    `json:"intf"`
	Uintf   uint   `json:"uintf"`
	Stringf string `json:"stringf"`
	Uuid    string `json:"uuid"`
	Ip      string `json:"ip"`
	Email   string `json:"email"`
	Int32f  int32  `json:"int32f"`
	Uint32f uint32 `json:"uint32f"`
	Int64f  int64  `json:"int64f"`
	Uint64f uint64 `json:"uint64f"`
}

var tenFieldsStructWithTagResult tenFieldsStructWithTag

func init() {
	var err error
	err = json.Unmarshal(tenFieldsByte, &tenFieldsByteMap)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(tenFieldsByte, &tenFieldsStructWoTagResult)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(tenFieldsByte, &tenFieldsStructWithTagResult)
	if err != nil {
		panic(err)
	}
}
