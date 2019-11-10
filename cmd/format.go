package main

import (
	"strings"
	"fmt"
)

var s = `
Benchmark_10Fields_Iterator_Object/jsoniter-readObj-8            1215607               917 ns/op             144 B/op         14 allocs/op
Benchmark_10Fields_Iterator_Object/jsoniter-readObjCB-8          1254342               934 ns/op             144 B/op         14 allocs/op
Benchmark_10Fields_Iterator_Object/jsonparser-8                  1274337               946 ns/op              80 B/op          4 allocs/op
Benchmark_10Fields_Unmarshal_Interface/json-8                     255212              4843 ns/op            1414 B/op         36 allocs/op
Benchmark_10Fields_Unmarshal_Interface/jsoniter-8                 342501              3603 ns/op            1350 B/op         38 allocs/op
Benchmark_10Fields_Unmarshal_Interface/jsoniter-compatible-8      333117              3632 ns/op            1350 B/op         38 allocs/op
Benchmark_10Fields_Unmarshal_Interface/djson-8                    521079              2276 ns/op            1174 B/op         27 allocs/op
Benchmark_10Fields_Unmarshal_Interface/easyjson-8                 428066              2967 ns/op            1173 B/op         27 allocs/op
Benchmark_10Fields_Unmarshal_Interface/ujson-8                    352647              3348 ns/op            1494 B/op         38 allocs/op
Benchmark_10Fields_Unmarshal_Interface/ugorji-8                   230550              5274 ns/op            2222 B/op         36 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/json-8                 272455              4720 ns/op             432 B/op         14 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/jsoniter-8            1000000              1052 ns/op             192 B/op          5 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/jsoniter-compatible-8                 1000000              1079 ns/op             192 B/op          5 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/ugorji-8                               521164              2567 ns/op             832 B/op          7 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/json-8                                   255205              4921 ns/op             432 B/op         14 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/jsoniter-8                               799615              1551 ns/op             256 B/op         15 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/jsoniter-compatible-8                    799477              1575 ns/op             256 B/op         15 allocs/op
Benchmark_20Fields_Unmarshal_Interface/json-8                                     124928              9957 ns/op            3003 B/op         67 allocs/op
Benchmark_20Fields_Unmarshal_Interface/jsoniter-8                                 155706              8053 ns/op            3051 B/op         73 allocs/op
Benchmark_20Fields_Unmarshal_Interface/jsoniter-compatible-8                      159856              7700 ns/op            3053 B/op         73 allocs/op
Benchmark_20Fields_Unmarshal_Interface/djson-8                                    239824              5003 ns/op            2715 B/op         52 allocs/op
Benchmark_20Fields_Unmarshal_Interface/easyjson-8                                 181693              6367 ns/op            2716 B/op         52 allocs/op
Benchmark_20Fields_Unmarshal_Interface/ujson-8                                    164239              7531 ns/op            3339 B/op         73 allocs/op
Benchmark_20Fields_Unmarshal_Interface/ugorji-8                                   119896             10558 ns/op            3763 B/op         61 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/json-8                                 137806              8765 ns/op             648 B/op         24 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/jsoniter-8                             428256              2872 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/jsoniter-compatible-8                  428274              2823 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/ugorji-8                               499628              2335 ns/op             832 B/op          7 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/json-8                                   118742              9936 ns/op             648 B/op         24 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/jsoniter-8                               413244              3012 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/jsoniter-compatible-8                    406504              2944 ns/op             512 B/op         29 allocs/op
`

func main() {
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		arr := strings.Split(line, " ")
		for _, elem := range arr {
			switch elem {
			case "", "ns/op", "B/op", "allocs/op":
			default:
				fmt.Printf("| %s ", elem)
			}
		}
		fmt.Println("|")
	}
}
