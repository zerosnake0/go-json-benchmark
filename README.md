# Introduction

This is a benchmark for the json packages.
You are welcome to open an issue if you find anything wrong or out of date.

# TL;DR

In conclusion, the `github.com/json-iterator/go` is your first choice,
it provides you with lots of useful features, oo design and high performance.

# How to run

```shell
go test -v -bench ./...
```

# Example result

|   |   |   | ns/op | B/op | allocs/op |
| - | - | - | ----- | ---- | --------- |
| 10Fields_Iterator_Object | jsoniter-readObj-8 | 1243933 | 1072 | 144 | 14 |
| 10Fields_Iterator_Object | jsoniter-readObjCB-8 | 1000000 | 1126 | 144 | 14 |
| 10Fields_Iterator_Object | jsonparser-8 | 856482 | 1171 | 80 | 4 |
| 10Fields_Iterator_Object | jzon-readObjCB-8 | 856426 | 1320 | 144 | 14 |
| 10Fields_Iterator_Object | jzon-readObj-8 | 749488 | 1386 | 144 | 14 |
|   |   |   |   |   |   |
| 10Fields_Unmarshal_Interface | djson-8 | 521389 | 2464 | 1174 | 27 |
| 10Fields_Unmarshal_Interface | easyjson-8 | 428254 | 2979 | 1174 | 27 |
| 10Fields_Unmarshal_Interface | jzon-fast-8 | 374584 | 3080 | 1190 | 28 |
| 10Fields_Unmarshal_Interface | jzon-8 | 333105 | 3266 | 1190 | 28 |
| 10Fields_Unmarshal_Interface | ujson-8 | 333092 | 3515 | 1494 | 38 |
| 10Fields_Unmarshal_Interface | jsoniter-compatible-8 | 333104 | 3908 | 1350 | 38 |
| 10Fields_Unmarshal_Interface | jsoniter-8 | 333102 | 4008 | 1350 | 38 |
| 10Fields_Unmarshal_Interface | ugorji-8 | 223449 | 5316 | 2222 | 36 |
| 10Fields_Unmarshal_Interface | json-8 | 235119 | 5346 | 1414 | 36 |
|   |   |   |   |   |   |
| 10Fields_Unmarshal_StructWithTag | jsoniter-8 | 1000000 | 1125 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | jsoniter-compatible-8 | 1000000 | 1125 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | jzon-8 | 749409 | 1679 | 212 | 15 |
| 10Fields_Unmarshal_StructWithTag | jzon-fast-8 | 750032 | 1757 | 212 | 15 |
| 10Fields_Unmarshal_StructWithTag | ugorji-8 | 444151 | 2584 | 832 | 7 |
| 10Fields_Unmarshal_StructWithTag | json-8 | 249820 | 4743 | 432 | 14 |
|   |   |   |   |   |   |
| 10Fields_Unmarshal_StructWoTag | jsoniter-8 | 705384 | 1539 | 256 | 15 |
| 10Fields_Unmarshal_StructWoTag | jsoniter-compatible-8 | 749451 | 1605 | 256 | 15 |
| 10Fields_Unmarshal_StructWoTag | jzon-8 | 571058 | 1978 | 212 | 15 |
| 10Fields_Unmarshal_StructWoTag | json-8 | 235126 | 5082 | 432 | 14 |
|   |   |   |   |   |   |
| 20Fields_Unmarshal_Interface | djson-8 | 213811 | 4961 | 2715 | 52 |
| 20Fields_Unmarshal_Interface | jzon-fast-8 | 181683 | 6637 | 2732 | 53 |
| 20Fields_Unmarshal_Interface | easyjson-8 | 187364 | 6815 | 2715 | 52 |
| 20Fields_Unmarshal_Interface | jzon-8 | 181689 | 6973 | 2732 | 53 |
| 20Fields_Unmarshal_Interface | ujson-8 | 159886 | 7523 | 3339 | 73 |
| 20Fields_Unmarshal_Interface | jsoniter-8 | 157783 | 7782 | 3052 | 73 |
| 20Fields_Unmarshal_Interface | jsoniter-compatible-8 | 151790 | 8056 | 3051 | 73 |
| 20Fields_Unmarshal_Interface | json-8 | 115303 | 10319 | 3003 | 67 |
| 20Fields_Unmarshal_Interface | ugorji-8 | 115303 | 10423 | 3763 | 61 |
|   |   |   |   |   |   |
| 20Fields_Unmarshal_StructWithTag | ugorji-8 | 499627 | 2381 | 832 | 7 |
| 20Fields_Unmarshal_StructWithTag | jsoniter-8 | 399579 | 2973 | 512 | 29 |
| 20Fields_Unmarshal_StructWithTag | jsoniter-compatible-8 | 413475 | 3016 | 512 | 29 |
| 20Fields_Unmarshal_StructWithTag | jzon-fast-8 | 370143 | 3163 | 408 | 29 |
| 20Fields_Unmarshal_StructWithTag | jzon-8 | 352695 | 3235 | 408 | 29 |
| 20Fields_Unmarshal_StructWithTag | json-8 | 131776 | 9926 | 648 | 24 |
|   |   |   |   |   |   |
| 20Fields_Unmarshal_StructWoTag | jsoniter-8 | 413481 | 2897 | 512 | 29 |
| 20Fields_Unmarshal_StructWoTag | jsoniter-compatible-8 | 413498 | 2931 | 512 | 29 |
| 20Fields_Unmarshal_StructWoTag | jzon-8 | 244722 | 4678 | 408 | 29 |
| 20Fields_Unmarshal_StructWoTag | json-8 | 111020 | 11700 | 648 | 24 |
|   |   |   |   |   |   |

# TODO

- add medium & large payload tests

# Packages tested

## feature list

| repo | unmarshal (interface) | unmarshal (struct) | iterator |
| -------------------------------------- | --------------- | --------------- | - |
| encoding/json                          | x               | x               |   |
| github.com/json-iterator/go            | x               | x               | x |
| github.com/buger/jsonparser            |                 |                 | x |
| github.com/a8m/djson                   | x               |                 |   |
| github.com/mailru/easyjson             | x               | <sup>(1)</sup>  |   |
| github.com/mreiferson/go-ujson         | x<sup>(2)</sup> |                 |   |
| github.com/ugorji/go/codec             | x<sup>(3)</sup> | x<sup>(4)</sup> |   |
| github.com/zerosnake0/jzon             | x               | x               | x |

### notes:
1. easyjson use its own marshaler/unmarshaler, which is not tested
2. ujson use ujson.numeric which is its internal type for numbers
3. ugorji need to be configured in order to be compatible with standard library
4. ugorji decode into structure:
   - can decode into a structure with tag
   - cannot decode into a structure without a tag unless all field names are as same as tag names
   - use its own marshaler/unmarshaler (codec.Selfer) interface to decode structure


# The following packages are not tested

| repo | reason |
| ------------------------------- | ------------------------------------------ |
| github.com/Jeffail/gabs         | the encoding/decoding uses `encoding/json` |
| github.com/bitly/go-simplejson  | the encoding/decoding uses `encoding/json` |
| github.com/antonholmquist/jason | the encoding/decoding uses `encoding/json` |
| github.com/pquerna/ffjson       | the encoding/decoding uses `encoding/json` (if the custom interface is not provided) |
