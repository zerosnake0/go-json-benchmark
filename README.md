[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

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

|     |     |     | ns/op | B/op | allocs/op |
| --- | --- | --- | ----- | ---- | --------- |
| 10Fields_Iterator_Object | jsoniter-readObjCB-8 | 1458022 | 810 | 144 | 14 |
| 10Fields_Iterator_Object | jsoniter-readObj-8 | 1369156 | 813 | 144 | 14 |
| 10Fields_Iterator_Object | jsonparser-8 | 1357131 | 908 | 80 | 4 |
| 10Fields_Iterator_Object | jzon-readObj-8 | 1239000 | 955 | 144 | 14 |
| 10Fields_Iterator_Object | jzon-readObjCB-8 | 1204593 | 991 | 144 | 14 |
|     |     |     |       |      |           |
| 10Fields_Unmarshal_Interface | djson-8 | 440643 | 2576 | 1174 | 27 |
| 10Fields_Unmarshal_Interface | easyjson-8 | 422460 | 2894 | 1174 | 27 |
| 10Fields_Unmarshal_Interface | jzon-8 | 414544 | 2904 | 1190 | 28 |
| 10Fields_Unmarshal_Interface | jzon-fast-8 | 425950 | 2917 | 1190 | 28 |
| 10Fields_Unmarshal_Interface | ujson-8 | 337881 | 3494 | 1494 | 38 |
| 10Fields_Unmarshal_Interface | jsoniter-compatible-8 | 320720 | 3758 | 1350 |
38 |
| 10Fields_Unmarshal_Interface | jsoniter-8 | 347126 | 4216 | 1350 | 38 |
| 10Fields_Unmarshal_Interface | json-8 | 229218 | 5532 | 1414 | 36 |
| 10Fields_Unmarshal_Interface | ugorji-8 | 212991 | 5608 | 2478 | 36 |
|     |     |     |       |      |           |
| 10Fields_Unmarshal_StructWithTag | jsoniter-8 | 1000000 | 1028 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | jsoniter-compatible-8 | 1223282 | 1034 | 19
2 | 5 |
| 10Fields_Unmarshal_StructWithTag | jzon-8 | 1000000 | 1106 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | jzon-fast-8 | 1000000 | 1132 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | ugorji-8 | 476871 | 2705 | 1088 | 7 |
| 10Fields_Unmarshal_StructWithTag | json-8 | 274917 | 4346 | 432 | 14 |
|     |     |     |       |      |           |
| 10Fields_Unmarshal_StructWoTag | jsoniter-compatible-8 | 887858 | 1406 | 256 |
 15 |
| 10Fields_Unmarshal_StructWoTag | jsoniter-8 | 872038 | 1451 | 256 | 15 |
| 10Fields_Unmarshal_StructWoTag | jzon-8 | 694782 | 1607 | 192 | 5 |
| 10Fields_Unmarshal_StructWoTag | json-8 | 242374 | 4879 | 432 | 14 |
|     |     |     |       |      |           |
| 10Fields_Valid | gjson-8 | 2991547 | 400 | 0 | 0 |
| 10Fields_Valid | jzon-8 | 2099695 | 567 | 0 | 0 |
| 10Fields_Valid | jsoniter-compatible-8 | 1583869 | 736 | 64 | 10 |
| 10Fields_Valid | jsoniter-8 | 1628780 | 758 | 64 | 10 |
| 10Fields_Valid | json-8 | 1000000 | 1114 | 0 | 0 |
|     |     |     |       |      |           |
| 20Fields_Unmarshal_Interface | jzon-fast-8 | 187482 | 6285 | 2732 | 53 |
| 20Fields_Unmarshal_Interface | jzon-8 | 193900 | 6516 | 2731 | 53 |
| 20Fields_Unmarshal_Interface | djson-8 | 238552 | 6606 | 2716 | 52 |
| 20Fields_Unmarshal_Interface | easyjson-8 | 194299 | 6974 | 2716 | 52 |
| 20Fields_Unmarshal_Interface | ujson-8 | 163780 | 7551 | 3340 | 73 |
| 20Fields_Unmarshal_Interface | jsoniter-compatible-8 | 150829 | 7872 | 3051 |
73 |
| 20Fields_Unmarshal_Interface | jsoniter-8 | 139962 | 8286 | 3051 | 73 |
| 20Fields_Unmarshal_Interface | json-8 | 121875 | 10689 | 3003 | 67 |
| 20Fields_Unmarshal_Interface | ugorji-8 | 119270 | 10768 | 4018 | 61 |
|     |     |     |       |      |           |
| 20Fields_Unmarshal_StructWithTag | jzon-8 | 512673 | 2345 | 368 | 9 |
| 20Fields_Unmarshal_StructWithTag | jzon-fast-8 | 497636 | 2402 | 368 | 9 |
| 20Fields_Unmarshal_StructWithTag | ugorji-8 | 484503 | 2532 | 1088 | 7 |
| 20Fields_Unmarshal_StructWithTag | jsoniter-8 | 420992 | 2780 | 512 | 29 |
| 20Fields_Unmarshal_StructWithTag | jsoniter-compatible-8 | 420163 | 2877 | 512
 | 29 |
| 20Fields_Unmarshal_StructWithTag | json-8 | 145645 | 8466 | 648 | 24 |
|     |     |     |       |      |           |
| 20Fields_Unmarshal_StructWoTag | jsoniter-compatible-8 | 425262 | 2724 | 512 |
 29 |
| 20Fields_Unmarshal_StructWoTag | jsoniter-8 | 435942 | 2758 | 512 | 29 |
| 20Fields_Unmarshal_StructWoTag | jzon-8 | 251348 | 4908 | 1057 | 9 |
| 20Fields_Unmarshal_StructWoTag | json-8 | 116900 | 9759 | 648 | 24 |
|     |     |     |       |      |           |
| 20Fields_Valid | gjson-8 | 1483011 | 776 | 0 | 0 |
| 20Fields_Valid | jzon-8 | 1000000 | 1070 | 0 | 0 |
| 20Fields_Valid | jsoniter-compatible-8 | 813631 | 1376 | 144 | 20 |
| 20Fields_Valid | jsoniter-8 | 772488 | 1418 | 144 | 20 |
| 20Fields_Valid | json-8 | 564940 | 2138 | 0 | 0 |
|     |     |     |       |      |           |

# TODO

- add medium & large payload tests
- auto benchmark

# Packages tested

## feature list

| repo | version | unmarshal (interface) | unmarshal (struct) | iterator | valid |
| -------------------------------------- | ------- | --------------- | --------------- | --- | --- |
| encoding/json                          | 1.15.2  | x               | x               |     |  x  |
| github.com/json-iterator/go            | 1.1.10  | x               | x               |  x  |  x  |
| github.com/buger/jsonparser            | 1.0.0   |                 |                 |  x  |     |
| github.com/a8m/djson                   | c02c5ae | x               |                 |     |     |
| github.com/mailru/easyjson             | 0.7.6   | x               | <sup>(1)</sup>  |     |     |
| github.com/mreiferson/go-ujson         | c02629f | x<sup>(2)</sup> |                 |     |     |
| github.com/ugorji/go/codec             | 1.1.10  | x<sup>(3)</sup> | x<sup>(4)</sup> |     |     |
| github.com/tidwall/gjson<sup>(5)</sup> | 1.6.1   |                 | <sup>(6)</sup>  |     |  x  |
| github.com/zerosnake0/jzon             | 0.0.5   | x               | x               |  x  |  x  |

### notes:
1. easyjson use its own marshaler/unmarshaler, which is not tested
2. ujson use ujson.numeric which is its internal type for numbers
3. ugorji need to be configured in order to be compatible with standard library
4. ugorji decode into structure:
   - can decode into a structure with tag
   - cannot decode into a structure without a tag unless all field names are as same as tag names
   - use its own marshaler/unmarshaler (codec.Selfer) interface to decode structure
5. gjson deprecated its unmarshal methods from `v1.4.0`
6. gjson only decode into structure fields with tag (`<= v1.3.6`)


# The following packages are not tested

| repo | reason |
| ------------------------------- | ------------------------------------------ |
| github.com/Jeffail/gabs         | the encoding/decoding uses `encoding/json` |
| github.com/bitly/go-simplejson  | the encoding/decoding uses `encoding/json` |
| github.com/antonholmquist/jason | the encoding/decoding uses `encoding/json` |
| github.com/pquerna/ffjson       | the encoding/decoding uses `encoding/json` (if the custom interface is not provided) |
