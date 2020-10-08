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
| 10Fields_Iterator_Object | jsoniter-readObj-8 | 1277052 | 911 | 144 | 14 |
| 10Fields_Iterator_Object | jsoniter-readObjCB-8 | 1319198 | 916 | 144 | 14 |
| 10Fields_Iterator_Object | jsonparser-8 | 1292194 | 928 | 80 | 4 |
| 10Fields_Iterator_Object | jzon-readObj-8 | 1000000 | 1030 | 144 | 14 |
| 10Fields_Iterator_Object | jzon-readObjCB-8 | 1000000 | 1045 | 144 | 14 |
|     |     |     |       |      |           |
| 10Fields_Unmarshal_Interface | djson-8 | 571047 | 2406 | 1174 | 27 |
| 10Fields_Unmarshal_Interface | jzon-8 | 386806 | 2843 | 1190 | 28 |
| 10Fields_Unmarshal_Interface | jzon-fast-8 | 444157 | 2870 | 1190 | 28 |
| 10Fields_Unmarshal_Interface | easyjson-8 | 428254 | 2919 | 1174 | 27 |
| 10Fields_Unmarshal_Interface | ujson-8 | 352842 | 3483 | 1494 | 38 |
| 10Fields_Unmarshal_Interface | jsoniter-compatible-8 | 342619 | 3663 | 1350 |
38 |
| 10Fields_Unmarshal_Interface | jsoniter-8 | 323967 | 3784 | 1350 | 38 |
| 10Fields_Unmarshal_Interface | json-8 | 244658 | 5056 | 1414 | 36 |
| 10Fields_Unmarshal_Interface | ugorji-8 | 226249 | 5277 | 2478 | 36 |
|     |     |     |       |      |           |
| 10Fields_Unmarshal_StructWithTag | jsoniter-compatible-8 | 998194 | 1039 | 192
 | 5 |
| 10Fields_Unmarshal_StructWithTag | jsoniter-8 | 1000000 | 1061 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | jzon-8 | 1000000 | 1121 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | jzon-fast-8 | 1000000 | 1141 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | ugorji-8 | 521378 | 2524 | 1088 | 7 |
| 10Fields_Unmarshal_StructWithTag | json-8 | 266540 | 4644 | 432 | 14 |
|     |     |     |       |      |           |
| 10Fields_Unmarshal_StructWoTag | jsoniter-compatible-8 | 799413 | 1528 | 256 |
 15 |
| 10Fields_Unmarshal_StructWoTag | jzon-8 | 705367 | 1588 | 192 | 5 |
| 10Fields_Unmarshal_StructWoTag | jsoniter-8 | 704836 | 1589 | 256 | 15 |
| 10Fields_Unmarshal_StructWoTag | json-8 | 235128 | 5094 | 432 | 14 |
|     |     |     |       |      |           |
| 10Fields_Valid | gjson-8 | 2337526 | 465 | 0 | 0 |
| 10Fields_Valid | jzon-8 | 1825243 | 644 | 0 | 0 |
| 10Fields_Valid | jsoniter-8 | 1266261 | 907 | 64 | 10 |
| 10Fields_Valid | jsoniter-compatible-8 | 1297790 | 936 | 64 | 10 |
| 10Fields_Valid | json-8 | 999217 | 1210 | 0 | 0 |
|     |     |     |       |      |           |
| 20Fields_Unmarshal_Interface | djson-8 | 226312 | 4952 | 2715 | 52 |
| 20Fields_Unmarshal_Interface | jzon-8 | 203245 | 6081 | 2731 | 53 |
| 20Fields_Unmarshal_Interface | jzon-fast-8 | 190340 | 6172 | 2732 | 53 |
| 20Fields_Unmarshal_Interface | easyjson-8 | 187369 | 6297 | 2716 | 52 |
| 20Fields_Unmarshal_Interface | ujson-8 | 159884 | 7655 | 3340 | 73 |
| 20Fields_Unmarshal_Interface | jsoniter-8 | 148044 | 8051 | 3052 | 73 |
| 20Fields_Unmarshal_Interface | jsoniter-compatible-8 | 149889 | 8105 | 3052 |
73 |
| 20Fields_Unmarshal_Interface | json-8 | 122362 | 9822 | 3003 | 67 |
| 20Fields_Unmarshal_Interface | ugorji-8 | 113108 | 10298 | 4020 | 61 |
|     |     |     |       |      |           |
| 20Fields_Unmarshal_StructWithTag | ugorji-8 | 460993 | 2483 | 1088 | 7 |
| 20Fields_Unmarshal_StructWithTag | jzon-8 | 386822 | 2613 | 368 | 9 |
| 20Fields_Unmarshal_StructWithTag | jzon-fast-8 | 479882 | 2630 | 368 | 9 |
| 20Fields_Unmarshal_StructWithTag | jsoniter-compatible-8 | 386816 | 2998 | 512
 | 29 |
| 20Fields_Unmarshal_StructWithTag | jsoniter-8 | 413494 | 3030 | 512 | 29 |
| 20Fields_Unmarshal_StructWithTag | json-8 | 130340 | 9075 | 648 | 24 |
|     |     |     |       |      |           |
| 20Fields_Unmarshal_StructWoTag | jsoniter-compatible-8 | 363232 | 3014 | 512 |
 29 |
| 20Fields_Unmarshal_StructWoTag | jsoniter-8 | 399932 | 3050 | 512 | 29 |
| 20Fields_Unmarshal_StructWoTag | jzon-8 | 218024 | 5182 | 1041 | 9 |
| 20Fields_Unmarshal_StructWoTag | json-8 | 116406 | 10325 | 648 | 24 |
|     |     |     |       |      |           |
| 20Fields_Valid | gjson-8 | 1356513 | 890 | 0 | 0 |
| 20Fields_Valid | jzon-8 | 999325 | 1220 | 0 | 0 |
| 20Fields_Valid | jsoniter-8 | 798775 | 1555 | 144 | 20 |
| 20Fields_Valid | jsoniter-compatible-8 | 856536 | 1573 | 144 | 20 |
| 20Fields_Valid | json-8 | 521455 | 2426 | 0 | 0 |
|     |     |     |       |      |           |

# TODO

- add medium & large payload tests
- auto benchmark

# Packages tested

## feature list

| repo | unmarshal (interface) | unmarshal (struct) | iterator | valid |
| -------------------------------------- | --------------- | --------------- | --- | --- |
| encoding/json                          | x               | x               |     |  x  |
| github.com/json-iterator/go            | x               | x               |  x  |  x  |
| github.com/buger/jsonparser            |                 |                 |  x  |     |
| github.com/a8m/djson                   | x               |                 |     |     |
| github.com/mailru/easyjson             | x               | <sup>(1)</sup>  |     |     |
| github.com/mreiferson/go-ujson         | x<sup>(2)</sup> |                 |     |     |
| github.com/ugorji/go/codec             | x<sup>(3)</sup> | x<sup>(4)</sup> |     |     |
| github.com/tidwall/gjson<sup>(5)</sup> |                 | <sup>(6)</sup>  |     |  x  |
| github.com/zerosnake0/jzon             | x               | x               |  x  |  x  |

### notes:
1. easyjson use its own marshaler/unmarshaler, which is not tested
2. ujson use ujson.numeric which is its internal type for numbers
3. ugorji need to be configured in order to be compatible with standard library
4. ugorji decode into structure:
   - can decode into a structure with tag
   - cannot decode into a structure without a tag unless all field names are as same as tag names
   - use its own marshaler/unmarshaler (codec.Selfer) interface to decode structure
5. gjson deprecated its unmarshal method from `v1.4.0`
6. gjson only decode into structure fields with tag (`<= v1.3.6`)


# The following packages are not tested

| repo | reason |
| ------------------------------- | ------------------------------------------ |
| github.com/Jeffail/gabs         | the encoding/decoding uses `encoding/json` |
| github.com/bitly/go-simplejson  | the encoding/decoding uses `encoding/json` |
| github.com/antonholmquist/jason | the encoding/decoding uses `encoding/json` |
| github.com/pquerna/ffjson       | the encoding/decoding uses `encoding/json` (if the custom interface is not provided) |
