# Packages tested

## feature list

| repo | unmarshal (interface) | unmarshal (struct) | iterator |
| -------------------------------------- |  - |  - |  - |
| encoding/json                          |  x |  x |  - |
| github.com/json-iterator/go            |  x |  x |  x |
| github.com/buger/jsonparser            |  - |  - |  x |
| github.com/a8m/djson                   |  x |  - |  - |
| github.com/mailru/easyjson             |  x | *1 |  x |
| github.com/mreiferson/go-ujson         | *2 |  - |  - |

### notes:
1. easyjson use its own marshaler/unmarshaler, which is not tested
2. ujson use ujson.numeric which is its internal type for numbers


# The packages not tested

| repo | reason |
| ------------------------------- | ------------------------------------------ |
| github.com/Jeffail/gabs         | the encoding/decoding uses `encoding/json` |
| github.com/bitly/go-simplejson  | the encoding/decoding uses `encoding/json` |
| github.com/antonholmquist/jason | the encoding/decoding uses `encoding/json` |