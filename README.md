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
| 10Fields_Iterator_Object | jsoniter-readObj-8 | 1316304 | 888 | 144 | 14 |
| 10Fields_Iterator_Object | jsoniter-readObjCB-8 | 1358041 | 898 | 144 | 14 |
| 10Fields_Iterator_Object | jsonparser-8 | 1227380 | 963 | 80 | 4 |
| 10Fields_Iterator_Object | jzon-readObj-8 | 1203966 | 993 | 144 | 14 |
| 10Fields_Iterator_Object | jzon-readObjCB-8 | 1000000 | 1010 | 144 | 14 |
|   |   |   |   |   |   |
| 10Fields_Unmarshal_Interface | djson-8 | 599166 | 2327 | 1174 | 27 |
| 10Fields_Unmarshal_Interface | jzon-fast-8 | 443917 | 2773 | 1190 | 28 |
| 10Fields_Unmarshal_Interface | jzon-8 | 428084 | 2812 | 1190 | 28 |
| 10Fields_Unmarshal_Interface | easyjson-8 | 428072 | 2927 | 1174 | 27 |
| 10Fields_Unmarshal_Interface | ujson-8 | 374734 | 3290 | 1494 | 38 |
| 10Fields_Unmarshal_Interface | jsoniter-8 | 307460 | 3525 | 1350 | 38 |
| 10Fields_Unmarshal_Interface | jsoniter-compatible-8 | 333093 | 3581 | 1350 | 38 |
| 10Fields_Unmarshal_Interface | json-8 | 260616 | 4834 | 1414 | 36 |
| 10Fields_Unmarshal_Interface | ugorji-8 | 244724 | 4960 | 2221 | 36 |
|   |   |   |   |   |   |
| 10Fields_Unmarshal_StructWithTag | jsoniter-compatible-8 | 1000000 | 1011 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | jsoniter-8 | 999300 | 1035 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | jzon-8 | 1000000 | 1073 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | jzon-fast-8 | 1000000 | 1088 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | ugorji-8 | 521652 | 2285 | 832 | 7 |
| 10Fields_Unmarshal_StructWithTag | json-8 | 272451 | 4488 | 432 | 14 |
|   |   |   |   |   |   |
| 10Fields_Unmarshal_StructWoTag | jsoniter-compatible-8 | 856506 | 1476 | 256 | 15 |
| 10Fields_Unmarshal_StructWoTag | jsoniter-8 | 800250 | 1484 | 256 | 15 |
| 10Fields_Unmarshal_StructWoTag | jzon-8 | 799386 | 1529 | 192 | 5 |
| 10Fields_Unmarshal_StructWoTag | json-8 | 249818 | 4807 | 432 | 14 |
|   |   |   |   |   |   |
| 20Fields_Unmarshal_Interface | djson-8 | 255135 | 4707 | 2716 | 52 |
| 20Fields_Unmarshal_Interface | jzon-fast-8 | 206701 | 5911 | 2732 | 53 |
| 20Fields_Unmarshal_Interface | jzon-8 | 199820 | 5925 | 2732 | 53 |
| 20Fields_Unmarshal_Interface | easyjson-8 | 193371 | 6174 | 2716 | 52 |
| 20Fields_Unmarshal_Interface | ujson-8 | 171274 | 7111 | 3340 | 73 |
| 20Fields_Unmarshal_Interface | jsoniter-compatible-8 | 157784 | 7427 | 3051 | 73 |
| 20Fields_Unmarshal_Interface | jsoniter-8 | 162044 | 7528 | 3051 | 73 |
| 20Fields_Unmarshal_Interface | json-8 | 114210 | 9516 | 3004 | 67 |
| 20Fields_Unmarshal_Interface | ugorji-8 | 119914 | 9847 | 3764 | 61 |
|   |   |   |   |   |   |
| 20Fields_Unmarshal_StructWithTag | ugorji-8 | 544728 | 2298 | 832 | 7 |
| 20Fields_Unmarshal_StructWithTag | jzon-8 | 521070 | 2381 | 368 | 9 |
| 20Fields_Unmarshal_StructWithTag | jzon-fast-8 | 499662 | 2425 | 368 | 9 |
| 20Fields_Unmarshal_StructWithTag | jsoniter-compatible-8 | 428072 | 2730 | 512 | 29 |
| 20Fields_Unmarshal_StructWithTag | jsoniter-8 | 428254 | 2816 | 512 | 29 |
| 20Fields_Unmarshal_StructWithTag | json-8 | 137854 | 8783 | 648 | 24 |
|   |   |   |   |   |   |
| 20Fields_Unmarshal_StructWoTag | jsoniter-8 | 444157 | 2758 | 512 | 29 |
| 20Fields_Unmarshal_StructWoTag | jsoniter-compatible-8 | 443912 | 2852 | 512 | 29 |
| 20Fields_Unmarshal_StructWoTag | jzon-8 | 244792 | 4963 | 1075 | 9 |
| 20Fields_Unmarshal_StructWoTag | json-8 | 119913 | 10048 | 648 | 24 |
|   |   |   |   |   |   |

# TODO

- add medium & large payload tests
- auto benchmark

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
