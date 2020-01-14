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

```shell
GOGC=off go test -bench=. -benchmem
```

The result may vary according to the different environment.
Use the result only as a reference.

|   |   |   | ns/op | B/op | allocs/op |
| - | - | - | ----- | ---- | --------- |
| 10Fields_Iterator_Object | jsoniter-readObjCB-8 | 1325042 | 934 | 144 | 14 |
| 10Fields_Iterator_Object | jsoniter-readObj-8 | 1202773 | 940 | 144 | 14 |
| 10Fields_Iterator_Object | jsonparser-8 | 1247815 | 998 | 80 | 4 |
| 10Fields_Iterator_Object | jzon-readObj-8 | 1000000 | 1066 | 144 | 14 |
| 10Fields_Iterator_Object | jzon-readObjCB-8 | 1000000 | 1090 | 144 | 14 |
|   |   |   |   |   |   |
| 10Fields_Unmarshal_Interface | djson-8 | 570604 | 2436 | 1174 | 27 |
| 10Fields_Unmarshal_Interface | jzon-fast-8 | 386988 | 2932 | 1190 | 28 |
| 10Fields_Unmarshal_Interface | easyjson-8 | 413498 | 2984 | 1174 | 27 |
| 10Fields_Unmarshal_Interface | jzon-8 | 428096 | 3235 | 1190 | 28 |
| 10Fields_Unmarshal_Interface | ujson-8 | 352674 | 3416 | 1494 | 38 |
| 10Fields_Unmarshal_Interface | jsoniter-8 | 324094 | 3554 | 1350 | 38 |
| 10Fields_Unmarshal_Interface | jsoniter-compatible-8 | 352828 | 3562 | 1350 | 38 |
| 10Fields_Unmarshal_Interface | ugorji-8 | 239774 | 5138 | 2222 | 36 |
| 10Fields_Unmarshal_Interface | json-8 | 249818 | 5616 | 1414 | 36 |
|   |   |   |   |   |   |
| 10Fields_Unmarshal_StructWithTag | jsoniter-compatible-8 | 1000000 | 1149 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | jsoniter-8 | 922728 | 1180 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | jzon-8 | 922551 | 1286 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | jzon-fast-8 | 921544 | 1387 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | ugorji-8 | 521378 | 2466 | 832 | 7 |
| 10Fields_Unmarshal_StructWithTag | json-8 | 249806 | 4763 | 432 | 14 |
|   |   |   |   |   |   |
| 10Fields_Unmarshal_StructWoTag | jsoniter-8 | 799728 | 1510 | 256 | 15 |
| 10Fields_Unmarshal_StructWoTag | jsoniter-compatible-8 | 749427 | 1581 | 256 | 15 |
| 10Fields_Unmarshal_StructWoTag | jzon-8 | 521378 | 2424 | 457 | 5 |
| 10Fields_Unmarshal_StructWoTag | json-8 | 235059 | 4917 | 432 | 14 |
|   |   |   |   |   |   |
| 20Fields_Unmarshal_Interface | djson-8 | 249829 | 4911 | 2716 | 52 |
| 20Fields_Unmarshal_Interface | easyjson-8 | 193408 | 6173 | 2715 | 52 |
| 20Fields_Unmarshal_Interface | jzon-fast-8 | 190345 | 6619 | 2731 | 53 |
| 20Fields_Unmarshal_Interface | jzon-8 | 175712 | 6880 | 2731 | 53 |
| 20Fields_Unmarshal_Interface | ujson-8 | 164264 | 7292 | 3339 | 73 |
| 20Fields_Unmarshal_Interface | jsoniter-compatible-8 | 159880 | 7974 | 3051 | 73 |
| 20Fields_Unmarshal_Interface | jsoniter-8 | 151768 | 8091 | 3052 | 73 |
| 20Fields_Unmarshal_Interface | ugorji-8 | 121124 | 10336 | 3763 | 61 |
| 20Fields_Unmarshal_Interface | json-8 | 100768 | 10358 | 3004 | 67 |
|   |   |   |   |   |   |
| 20Fields_Unmarshal_StructWithTag | ugorji-8 | 521350 | 2292 | 832 | 7 |
| 20Fields_Unmarshal_StructWithTag | jzon-8 | 479347 | 2758 | 368 | 9 |
| 20Fields_Unmarshal_StructWithTag | jzon-fast-8 | 363368 | 2795 | 368 | 9 |
| 20Fields_Unmarshal_StructWithTag | jsoniter-compatible-8 | 324087 | 3103 | 512 | 29 |
| 20Fields_Unmarshal_StructWithTag | jsoniter-8 | 399728 | 3262 | 512 | 29 |
| 20Fields_Unmarshal_StructWithTag | json-8 | 111030 | 9319 | 648 | 24 |
|   |   |   |   |   |   |
| 20Fields_Unmarshal_StructWoTag | jsoniter-8 | 428280 | 2911 | 512 | 29 |
| 20Fields_Unmarshal_StructWoTag | jsoniter-compatible-8 | 386800 | 2936 | 512 | 29 |
| 20Fields_Unmarshal_StructWoTag | jzon-8 | 230370 | 5426 | 1120 | 9 |
| 20Fields_Unmarshal_StructWoTag | json-8 | 113130 | 10774 | 648 | 24 |
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
