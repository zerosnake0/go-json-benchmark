package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

var s = `
Benchmark_10Fields_Iterator_Object/jsoniter-readObj-8            1202773               940 ns/op             144 B/op         14 allocs/op
Benchmark_10Fields_Iterator_Object/jsoniter-readObjCB-8          1325042               934 ns/op             144 B/op         14 allocs/op
Benchmark_10Fields_Iterator_Object/jsonparser-8                  1247815               998 ns/op              80 B/op          4 allocs/op
Benchmark_10Fields_Iterator_Object/jzon-readObj-8                1000000              1066 ns/op             144 B/op         14 allocs/op
Benchmark_10Fields_Iterator_Object/jzon-readObjCB-8              1000000              1090 ns/op             144 B/op         14 allocs/op
Benchmark_10Fields_Unmarshal_Interface/json-8                     249818              5616 ns/op            1414 B/op         36 allocs/op
Benchmark_10Fields_Unmarshal_Interface/jsoniter-8                 324094              3554 ns/op            1350 B/op         38 allocs/op
Benchmark_10Fields_Unmarshal_Interface/jsoniter-compatible-8      352828              3562 ns/op            1350 B/op         38 allocs/op
Benchmark_10Fields_Unmarshal_Interface/djson-8                    570604              2436 ns/op            1174 B/op         27 allocs/op
Benchmark_10Fields_Unmarshal_Interface/easyjson-8                 413498              2984 ns/op            1174 B/op         27 allocs/op
Benchmark_10Fields_Unmarshal_Interface/ujson-8                    352674              3416 ns/op            1494 B/op         38 allocs/op
Benchmark_10Fields_Unmarshal_Interface/ugorji-8                   239774              5138 ns/op            2222 B/op         36 allocs/op
Benchmark_10Fields_Unmarshal_Interface/jzon-8                     428096              3235 ns/op            1190 B/op         28 allocs/op
Benchmark_10Fields_Unmarshal_Interface/jzon-fast-8                386988              2932 ns/op            1190 B/op         28 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/json-8                 249806              4763 ns/op             432 B/op         14 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/jsoniter-8             922728              1180 ns/op             192 B/op          5 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/jsoniter-compatible-8                 1000000              1149 ns/op             192 B/op          5 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/ugorji-8                               521378              2466 ns/op             832 B/op          7 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/jzon-8                                 922551              1286 ns/op             192 B/op          5 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/jzon-fast-8                            921544              1387 ns/op             192 B/op          5 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/json-8                                   235059              4917 ns/op             432 B/op         14 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/jsoniter-8                               799728              1510 ns/op             256 B/op         15 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/jsoniter-compatible-8                    749427              1581 ns/op             256 B/op         15 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/jzon-8                                   521378              2424 ns/op             457 B/op          5 allocs/op
Benchmark_20Fields_Unmarshal_Interface/json-8                                     100768             10358 ns/op            3004 B/op         67 allocs/op
Benchmark_20Fields_Unmarshal_Interface/jsoniter-8                                 151768              8091 ns/op            3052 B/op         73 allocs/op
Benchmark_20Fields_Unmarshal_Interface/jsoniter-compatible-8                      159880              7974 ns/op            3051 B/op         73 allocs/op
Benchmark_20Fields_Unmarshal_Interface/djson-8                                    249829              4911 ns/op            2716 B/op         52 allocs/op
Benchmark_20Fields_Unmarshal_Interface/easyjson-8                                 193408              6173 ns/op            2715 B/op         52 allocs/op
Benchmark_20Fields_Unmarshal_Interface/ujson-8                                    164264              7292 ns/op            3339 B/op         73 allocs/op
Benchmark_20Fields_Unmarshal_Interface/ugorji-8                                   121124             10336 ns/op            3763 B/op         61 allocs/op
Benchmark_20Fields_Unmarshal_Interface/jzon-8                                     175712              6880 ns/op            2731 B/op         53 allocs/op
Benchmark_20Fields_Unmarshal_Interface/jzon-fast-8                                190345              6619 ns/op            2731 B/op         53 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/json-8                                 111030              9319 ns/op             648 B/op         24 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/jsoniter-8                             399728              3262 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/jsoniter-compatible-8                  324087              3103 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/ugorji-8                               521350              2292 ns/op             832 B/op          7 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/jzon-8                                 479347              2758 ns/op             368 B/op          9 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/jzon-fast-8                            363368              2795 ns/op             368 B/op          9 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/json-8                                   113130             10774 ns/op             648 B/op         24 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/jsoniter-8                               428280              2911 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/jsoniter-compatible-8                    386800              2936 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/jzon-8                                   230370              5426 ns/op            1120 B/op          9 allocs/op
`

type result struct {
	Name     string
	N        int
	NsOp     int
	BOp      int
	AllocsOp int
}

type resList []result

var _ sort.Interface = resList{}

func (l resList) Len() int {
	return len(l)
}

func (l resList) Less(i, j int) bool {
	return l[i].NsOp < l[j].NsOp
}

func (l resList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

const (
	head  = "|   |   |   | ns/op | B/op | allocs/op |"
	head2 = "| - | - | - | ----- | ---- | --------- |"
	sep   = "|   |   |   |   |   |   |"
)

func main() {
	var keys []string
	m := map[string]resList{}
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		line = strings.TrimPrefix(line, "Benchmark_")
		var (
			name string
			res  result
		)
		n, err := fmt.Sscanf(line, "%s %d %d ns/op %d B/op %d allocs/op",
			&name, &res.N, &res.NsOp, &res.BOp, &res.AllocsOp)
		if err != nil {
			log.Fatal(err)
		}
		if n != 5 {
			log.Fatalf("%d parsed instead of 5", n)
		}

		arr := strings.SplitN(name, "/", 2)
		res.Name = arr[1]
		name = arr[0]
		if _, ok := m[name]; !ok {
			keys = append(keys, name)
		}
		m[name] = append(m[name], res)

	}
	fmt.Println(head)
	fmt.Println(head2)
	for _, k := range keys {
		arr := m[k]
		sort.Sort(arr)
		for _, item := range arr {
			fmt.Printf("| %s | %s | %d | %d | %d | %d |\n", k, item.Name,
				item.N, item.NsOp, item.BOp, item.AllocsOp)
		}
		fmt.Println(sep)
	}
}
