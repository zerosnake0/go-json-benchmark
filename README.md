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

|  |  | ns/op | B/op | allocs/op |
| - | - | - | - | - |
| Benchmark_10Fields_Iterator_Object/jsoniter-readObj-8 | 1215607 | 917 | 144 | 14 |
| Benchmark_10Fields_Iterator_Object/jsoniter-readObjCB-8 | 1254342 | 934 | 144 | 14 |
| Benchmark_10Fields_Iterator_Object/jsonparser-8 | 1274337 | 946 | 80 | 4 |
|
| Benchmark_10Fields_Unmarshal_Interface/json-8 | 255212 | 4843 | 1414 | 36 |
| Benchmark_10Fields_Unmarshal_Interface/jsoniter-8 | 342501 | 3603 | 1350 | 38 |
| Benchmark_10Fields_Unmarshal_Interface/jsoniter-compatible-8 | 333117 | 3632 | 1350 | 38 |
| Benchmark_10Fields_Unmarshal_Interface/djson-8 | 521079 | 2276 | 1174 | 27 |
| Benchmark_10Fields_Unmarshal_Interface/easyjson-8 | 428066 | 2967 | 1173 | 27 |
| Benchmark_10Fields_Unmarshal_Interface/ujson-8 | 352647 | 3348 | 1494 | 38 |
| Benchmark_10Fields_Unmarshal_Interface/ugorji-8 | 230550 | 5274 | 2222 | 36 |
|
| Benchmark_10Fields_Unmarshal_StructWithTag/json-8 | 272455 | 4720 | 432 | 14 |
| Benchmark_10Fields_Unmarshal_StructWithTag/jsoniter-8 | 1000000 | 1052 | 192 | 5 |
| Benchmark_10Fields_Unmarshal_StructWithTag/jsoniter-compatible-8 | 1000000 | 1079 | 192 | 5 |
| Benchmark_10Fields_Unmarshal_StructWithTag/ugorji-8 | 521164 | 2567 | 832 | 7 |
|
| Benchmark_10Fields_Unmarshal_StructWoTag/json-8 | 255205 | 4921 | 432 | 14 |
| Benchmark_10Fields_Unmarshal_StructWoTag/jsoniter-8 | 799615 | 1551 | 256 | 15 |
| Benchmark_10Fields_Unmarshal_StructWoTag/jsoniter-compatible-8 | 799477 | 1575 | 256 | 15 |
|
| Benchmark_20Fields_Unmarshal_Interface/json-8 | 124928 | 9957 | 3003 | 67 |
| Benchmark_20Fields_Unmarshal_Interface/jsoniter-8 | 155706 | 8053 | 3051 | 73 |
| Benchmark_20Fields_Unmarshal_Interface/jsoniter-compatible-8 | 159856 | 7700 | 3053 | 73 |
| Benchmark_20Fields_Unmarshal_Interface/djson-8 | 239824 | 5003 | 2715 | 52 |
| Benchmark_20Fields_Unmarshal_Interface/easyjson-8 | 181693 | 6367 | 2716 | 52 |
| Benchmark_20Fields_Unmarshal_Interface/ujson-8 | 164239 | 7531 | 3339 | 73 |
| Benchmark_20Fields_Unmarshal_Interface/ugorji-8 | 119896 | 10558 | 3763 | 61 |
|
| Benchmark_20Fields_Unmarshal_StructWithTag/json-8 | 137806 | 8765 | 648 | 24 |
| Benchmark_20Fields_Unmarshal_StructWithTag/jsoniter-8 | 428256 | 2872 | 512 | 29 |
| Benchmark_20Fields_Unmarshal_StructWithTag/jsoniter-compatible-8 | 428274 | 2823 | 512 | 29 |
| Benchmark_20Fields_Unmarshal_StructWithTag/ugorji-8 | 499628 | 2335 | 832 | 7 |
|
| Benchmark_20Fields_Unmarshal_StructWoTag/json-8 | 118742 | 9936 | 648 | 24 |
| Benchmark_20Fields_Unmarshal_StructWoTag/jsoniter-8 | 413244 | 3012 | 512 | 29 |
| Benchmark_20Fields_Unmarshal_StructWoTag/jsoniter-compatible-8 | 406504 | 2944 | 512 | 29 |

# TODO

- add medium & large payload tests

# Packages tested

## feature list

| repo | unmarshal (interface) | unmarshal (struct) | iterator |
| -------------------------------------- | -- | -- | -- |
| encoding/json                          |  x |  x |    |
| github.com/json-iterator/go            |  x |  x |  x |
| github.com/buger/jsonparser            |    |    |  x |
| github.com/a8m/djson                   |  x |    |    |
| github.com/mailru/easyjson             |  x | *1 |    |
| github.com/mreiferson/go-ujson         | *2 |    |    |
| github.com/ugorji/go/codec             | *3 | *4 |    |

### notes:
1. easyjson use its own marshaler/unmarshaler, which is not tested
2. ujson use ujson.numeric which is its internal type for numbers
3. ugorji need to be configured in order to be compatible with standard library
4. ugorji decode into structure:
   - can decode into a structure with tag
   - cannot decode into a structure without a tag unless all field names are as same as field name
   - use its own marshaler/unmarshaler (codec.Selfer) interface to decode structure


# The following packages are not tested

| repo | reason |
| ------------------------------- | ------------------------------------------ |
| github.com/Jeffail/gabs         | the encoding/decoding uses `encoding/json` |
| github.com/bitly/go-simplejson  | the encoding/decoding uses `encoding/json` |
| github.com/antonholmquist/jason | the encoding/decoding uses `encoding/json` |
| github.com/pquerna/ffjson       | the encoding/decoding uses `encoding/json` (if the custom interface is not provided) |