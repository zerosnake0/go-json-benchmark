# Introduction

This is a benchmark for the json packages.
You are welcome to open an issue if you find anything wrong or out of date.

# TL;DR

In conclusion, the `github.com/json-iterator/go` is your first choice,
it provides you with lots of useful features, oo design and high performance.

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