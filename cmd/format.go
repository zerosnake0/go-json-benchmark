package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

var s = `
Benchmark_10Fields_Iterator_Object/jsoniter-readObj-8            1362319               882 ns/op             144 B/op         14 allocs/op
Benchmark_10Fields_Iterator_Object/jsoniter-readObjCB-8          1342808               901 ns/op             144 B/op         14 allocs/op
Benchmark_10Fields_Iterator_Object/jsonparser-8                  1000000              1020 ns/op              80 B/op          4 allocs/op
Benchmark_10Fields_Iterator_Object/jzon-readObj-8                1168579              1017 ns/op             144 B/op         14 allocs/op
Benchmark_10Fields_Iterator_Object/jzon-readObjCB-8              1000000              1029 ns/op             144 B/op         14 allocs/op
Benchmark_10Fields_Unmarshal_Interface/json-8                     250053              4790 ns/op            1414 B/op         36 allocs/op
Benchmark_10Fields_Unmarshal_Interface/jsoniter-8                 342622              3569 ns/op            1350 B/op         38 allocs/op
Benchmark_10Fields_Unmarshal_Interface/jsoniter-compatible-8      342631              3587 ns/op            1350 B/op         38 allocs/op
Benchmark_10Fields_Unmarshal_Interface/djson-8                    545070              2269 ns/op            1174 B/op         27 allocs/op
Benchmark_10Fields_Unmarshal_Interface/easyjson-8                 412998              2862 ns/op            1174 B/op         27 allocs/op
Benchmark_10Fields_Unmarshal_Interface/ujson-8                    342634              3449 ns/op            1494 B/op         38 allocs/op
Benchmark_10Fields_Unmarshal_Interface/ugorji-8                   235129              5137 ns/op            2222 B/op         36 allocs/op
Benchmark_10Fields_Unmarshal_Interface/gjson-8                    386816              3037 ns/op            1318 B/op         16 allocs/op
Benchmark_10Fields_Unmarshal_Interface/jzon-8                     413323              2869 ns/op            1190 B/op         28 allocs/op
Benchmark_10Fields_Unmarshal_Interface/jzon-fast-8                413500              2856 ns/op            1190 B/op         28 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/json-8                 266504              4506 ns/op             432 B/op         14 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/jsoniter-8            1000000              1037 ns/op             192 B/op          5 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/jsoniter-compatible-8                 1000000              1081 ns/op             192 B/op          5 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/ugorji-8                               479446              2365 ns/op             832 B/op          7 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/gjson-8                                499610              2386 ns/op             480 B/op          3 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/jzon-8                                1000000              1125 ns/op             192 B/op          5 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/jzon-fast-8                           1000000              1101 ns/op             192 B/op          5 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/json-8                                   249825              4879 ns/op             432 B/op         14 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/jsoniter-8                               799434              1530 ns/op             256 B/op         15 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/jsoniter-compatible-8                    798753              1486 ns/op             256 B/op         15 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/jzon-8                                   749488              1558 ns/op             192 B/op          5 allocs/op
Benchmark_10Fields_Valid/json-8                                                   998218              1221 ns/op              72 B/op          2 allocs/op
Benchmark_10Fields_Valid/jsoniter-8                                              1476784               796 ns/op              64 B/op         10 allocs/op
Benchmark_10Fields_Valid/jsoniter-compatible-8                                   1487803               826 ns/op              64 B/op         10 allocs/op
Benchmark_10Fields_Valid/gjson-8                                                 2393536               437 ns/op               0 B/op          0 allocs/op
Benchmark_10Fields_Valid/jzon-8                                                  1921750               643 ns/op               0 B/op          0 allocs/op
Benchmark_20Fields_Unmarshal_Interface/json-8                                     123606              9828 ns/op            3003 B/op         67 allocs/op
Benchmark_20Fields_Unmarshal_Interface/jsoniter-8                                 157754              7676 ns/op            3052 B/op         73 allocs/op
Benchmark_20Fields_Unmarshal_Interface/jsoniter-compatible-8                      159882              7680 ns/op            3053 B/op         73 allocs/op
Benchmark_20Fields_Unmarshal_Interface/djson-8                                    210356              5138 ns/op            2716 B/op         52 allocs/op
Benchmark_20Fields_Unmarshal_Interface/easyjson-8                                 190302              6279 ns/op            2715 B/op         52 allocs/op
Benchmark_20Fields_Unmarshal_Interface/ujson-8                                    166578              7515 ns/op            3339 B/op         73 allocs/op
Benchmark_20Fields_Unmarshal_Interface/ugorji-8                                   118713             10090 ns/op            3763 B/op         61 allocs/op
Benchmark_20Fields_Unmarshal_Interface/gjson-8                                    181651              6567 ns/op            2923 B/op         27 allocs/op
Benchmark_20Fields_Unmarshal_Interface/jzon-8                                     196587              6180 ns/op            2732 B/op         53 allocs/op
Benchmark_20Fields_Unmarshal_Interface/jzon-fast-8                                187329              6191 ns/op            2732 B/op         53 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/json-8                                 137854              8849 ns/op             648 B/op         24 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/jsoniter-8                             428065              2840 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/jsoniter-compatible-8                  413689              2905 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/ugorji-8                               521600              2306 ns/op             832 B/op          7 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/gjson-8                                278880              4248 ns/op             896 B/op          3 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/jzon-8                                 499653              2546 ns/op             368 B/op          9 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/jzon-fast-8                            479487              2502 ns/op             368 B/op          9 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/json-8                                   117565             10138 ns/op             648 B/op         24 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/jsoniter-8                               428092              2878 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/jsoniter-compatible-8                    413500              2836 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/jzon-8                                   244732              5164 ns/op            1082 B/op          9 allocs/op
Benchmark_20Fields_Valid/json-8                                                   521661              2576 ns/op              72 B/op          2 allocs/op
Benchmark_20Fields_Valid/jsoniter-8                                               799386              1522 ns/op             144 B/op         20 allocs/op
Benchmark_20Fields_Valid/jsoniter-compatible-8                                    798657              1477 ns/op             144 B/op         20 allocs/op
Benchmark_20Fields_Valid/gjson-8                                                 1286642               863 ns/op               0 B/op          0 allocs/op
Benchmark_20Fields_Valid/jzon-8                                                  1052240              1223 ns/op               0 B/op          0 allocs/op
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
