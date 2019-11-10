package json_benchmark

import (
	"encoding/json"
)

var twentyFieldsByte = []byte(`{
	"intf": -123,
	"uintf": 456,
	"stringf": "abc",
	"uuidf": "de305d54-75b4-431b-adb2-eb6b9e546014",
	"ipf": "127.0.0.1",
	"emailf": "test@example.com",
	"int32f": -164281,
	"uint32f": 2938729,
	"int64f": -8973,
	"uint64f":89302174,
	"intf2": -123,
	"uintf2": 456,
	"stringf2": "abc",
	"uuidf2": "de305d54-75b4-431b-adb2-eb6b9e546014",
	"ipf2": "127.0.0.1",
	"emailf2": "test@example.com",
	"int32f2": -164281,
	"uint32f2": 2938729,
	"int64f2": -8973,
	"uint64f2":89302174
}`)

var twentyFieldsByteMap map[string]interface{}

type twentyFieldsStructWoTag struct {
	Intf     int
	Uintf    uint
	Stringf  string
	Uuidf    string
	Ipf      string
	Emailf   string
	Int32f   int32
	Uint32f  uint32
	Int64f   int64
	Uint64f  uint64
	Intf2    int
	Uintf2   uint
	Stringf2 string
	Uuidf2   string
	Ipf2     string
	Emailf2  string
	Int32f2  int32
	Uint32f2 uint32
	Int64f2  int64
	Uint64f2 uint64
}

var twentyFieldsStructWoTagResult twentyFieldsStructWoTag

type twentyFieldsStructWithTag struct {
	Intf     int    `json:"intf"`
	Uintf    uint   `json:"uint"`
	Stringf  string `json:"stringf"`
	Uuidf    string `json:"uuidf"`
	Ipf      string `json:"ipf"`
	Emailf   string `json:"emailf"`
	Int32f   int32  `json:"int32f"`
	Uint32f  uint32 `json:"uint32f"`
	Int64f   int64  `json:"int64f"`
	Uint64f  uint64 `json:"uint64f"`
	Intf2    int    `json:"intf2 "`
	Uintf2   uint   `json:"uint"`
	Stringf2 string `json:"stringf2 "`
	Uuidf2   string `json:"uuidf2 "`
	Ipf2     string `json:"ipf2 "`
	Emailf2  string `json:"emailf2 "`
	Int32f2  int32  `json:"int32f2 "`
	Uint32f2 uint32 `json:"uint32f2 "`
	Int64f2  int64  `json:"int64f2 "`
	Uint64f2 uint64 `json:"uint64f2 "`
}

var twentyFieldsStructWithTagResult twentyFieldsStructWithTag

func init() {
	var err error
	err = json.Unmarshal(twentyFieldsByte, &twentyFieldsByteMap)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(twentyFieldsByte, &twentyFieldsStructWoTagResult)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(twentyFieldsByte, &twentyFieldsStructWithTagResult)
	if err != nil {
		panic(err)
	}
}
