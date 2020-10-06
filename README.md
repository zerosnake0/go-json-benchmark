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
| 10Fields_Iterator_Object | jsoniter-readObj-8 | 1362319 | 882 | 144 | 14 |
| 10Fields_Iterator_Object | jsoniter-readObjCB-8 | 1342808 | 901 | 144 | 14 |
| 10Fields_Iterator_Object | jzon-readObj-8 | 1168579 | 1017 | 144 | 14 |
| 10Fields_Iterator_Object | jsonparser-8 | 1000000 | 1020 | 80 | 4 |
| 10Fields_Iterator_Object | jzon-readObjCB-8 | 1000000 | 1029 | 144 | 14 |
|     |     |     |       |      |           |
| 10Fields_Unmarshal_Interface | djson-8 | 545070 | 2269 | 1174 | 27 |
| 10Fields_Unmarshal_Interface | jzon-fast-8 | 413500 | 2856 | 1190 | 28 |
| 10Fields_Unmarshal_Interface | easyjson-8 | 412998 | 2862 | 1174 | 27 |
| 10Fields_Unmarshal_Interface | jzon-8 | 413323 | 2869 | 1190 | 28 |
| 10Fields_Unmarshal_Interface | gjson-8 | 386816 | 3037 | 1318 | 16 |
| 10Fields_Unmarshal_Interface | ujson-8 | 342634 | 3449 | 1494 | 38 |
| 10Fields_Unmarshal_Interface | jsoniter-8 | 342622 | 3569 | 1350 | 38 |
| 10Fields_Unmarshal_Interface | jsoniter-compatible-8 | 342631 | 3587 | 1350 | 38 |
| 10Fields_Unmarshal_Interface | json-8 | 250053 | 4790 | 1414 | 36 |
| 10Fields_Unmarshal_Interface | ugorji-8 | 235129 | 5137 | 2222 | 36 |
|     |     |     |       |      |           |
| 10Fields_Unmarshal_StructWithTag | jsoniter-8 | 1000000 | 1037 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | jsoniter-compatible-8 | 1000000 | 1081 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | jzon-fast-8 | 1000000 | 1101 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | jzon-8 | 1000000 | 1125 | 192 | 5 |
| 10Fields_Unmarshal_StructWithTag | ugorji-8 | 479446 | 2365 | 832 | 7 |
| 10Fields_Unmarshal_StructWithTag | gjson-8 | 499610 | 2386 | 480 | 3 |
| 10Fields_Unmarshal_StructWithTag | json-8 | 266504 | 4506 | 432 | 14 |
|     |     |     |       |      |           |
| 10Fields_Unmarshal_StructWoTag | jsoniter-compatible-8 | 798753 | 1486 | 256 | 15 |
| 10Fields_Unmarshal_StructWoTag | jsoniter-8 | 799434 | 1530 | 256 | 15 |
| 10Fields_Unmarshal_StructWoTag | jzon-8 | 749488 | 1558 | 192 | 5 |
| 10Fields_Unmarshal_StructWoTag | json-8 | 249825 | 4879 | 432 | 14 |
|     |     |     |       |      |           |
| 10Fields_Valid | gjson-8 | 2393536 | 437 | 0 | 0 |
| 10Fields_Valid | jzon-8 | 1921750 | 643 | 0 | 0 |
| 10Fields_Valid | jsoniter-8 | 1476784 | 796 | 64 | 10 |
| 10Fields_Valid | jsoniter-compatible-8 | 1487803 | 826 | 64 | 10 |
| 10Fields_Valid | json-8 | 998218 | 1221 | 72 | 2 |
|     |     |     |       |      |           |
| 20Fields_Unmarshal_Interface | djson-8 | 210356 | 5138 | 2716 | 52 |
| 20Fields_Unmarshal_Interface | jzon-8 | 196587 | 6180 | 2732 | 53 |
| 20Fields_Unmarshal_Interface | jzon-fast-8 | 187329 | 6191 | 2732 | 53 |
| 20Fields_Unmarshal_Interface | easyjson-8 | 190302 | 6279 | 2715 | 52 |
| 20Fields_Unmarshal_Interface | gjson-8 | 181651 | 6567 | 2923 | 27 |
| 20Fields_Unmarshal_Interface | ujson-8 | 166578 | 7515 | 3339 | 73 |
| 20Fields_Unmarshal_Interface | jsoniter-8 | 157754 | 7676 | 3052 | 73 |
| 20Fields_Unmarshal_Interface | jsoniter-compatible-8 | 159882 | 7680 | 3053 | 73 |
| 20Fields_Unmarshal_Interface | json-8 | 123606 | 9828 | 3003 | 67 |
| 20Fields_Unmarshal_Interface | ugorji-8 | 118713 | 10090 | 3763 | 61 |
|     |     |     |       |      |           |
| 20Fields_Unmarshal_StructWithTag | ugorji-8 | 521600 | 2306 | 832 | 7 |
| 20Fields_Unmarshal_StructWithTag | jzon-fast-8 | 479487 | 2502 | 368 | 9 |
| 20Fields_Unmarshal_StructWithTag | jzon-8 | 499653 | 2546 | 368 | 9 |
| 20Fields_Unmarshal_StructWithTag | jsoniter-8 | 428065 | 2840 | 512 | 29 |
| 20Fields_Unmarshal_StructWithTag | jsoniter-compatible-8 | 413689 | 2905 | 512 | 29 |
| 20Fields_Unmarshal_StructWithTag | gjson-8 | 278880 | 4248 | 896 | 3 |
| 20Fields_Unmarshal_StructWithTag | json-8 | 137854 | 8849 | 648 | 24 |
|     |     |     |       |      |           |
| 20Fields_Unmarshal_StructWoTag | jsoniter-compatible-8 | 413500 | 2836 | 512 | 29 |
| 20Fields_Unmarshal_StructWoTag | jsoniter-8 | 428092 | 2878 | 512 | 29 |
| 20Fields_Unmarshal_StructWoTag | jzon-8 | 244732 | 5164 | 1082 | 9 |
| 20Fields_Unmarshal_StructWoTag | json-8 | 117565 | 10138 | 648 | 24 |
|     |     |     |       |      |           |
| 20Fields_Valid | gjson-8 | 1286642 | 863 | 0 | 0 |
| 20Fields_Valid | jzon-8 | 1052240 | 1223 | 0 | 0 |
| 20Fields_Valid | jsoniter-compatible-8 | 798657 | 1477 | 144 | 20 |
| 20Fields_Valid | jsoniter-8 | 799386 | 1522 | 144 | 20 |
| 20Fields_Valid | json-8 | 521661 | 2576 | 72 | 2 |
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
| github.com/tidwall/gjson               | x               | x<sup>(5)</sup> |     |  x  |
| github.com/zerosnake0/jzon             | x               | x               |  x  |  x  |

### notes:
1. easyjson use its own marshaler/unmarshaler, which is not tested
2. ujson use ujson.numeric which is its internal type for numbers
3. ugorji need to be configured in order to be compatible with standard library
4. ugorji decode into structure:
   - can decode into a structure with tag
   - cannot decode into a structure without a tag unless all field names are as same as tag names
   - use its own marshaler/unmarshaler (codec.Selfer) interface to decode structure
5. gjson only decode into structure fields with tag


# The following packages are not tested

| repo | reason |
| ------------------------------- | ------------------------------------------ |
| github.com/Jeffail/gabs         | the encoding/decoding uses `encoding/json` |
| github.com/bitly/go-simplejson  | the encoding/decoding uses `encoding/json` |
| github.com/antonholmquist/jason | the encoding/decoding uses `encoding/json` |
| github.com/pquerna/ffjson       | the encoding/decoding uses `encoding/json` (if the custom interface is not provided) |
