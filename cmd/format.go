package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

var s = `
Benchmark_10Fields_Iterator_Object/jsoniter-readObj-8            1316304               888 ns/op             144 B/op         14 allocs/op
Benchmark_10Fields_Iterator_Object/jsoniter-readObjCB-8          1358041               898 ns/op             144 B/op         14 allocs/op
Benchmark_10Fields_Iterator_Object/jsonparser-8                  1227380               963 ns/op              80 B/op          4 allocs/op
Benchmark_10Fields_Iterator_Object/jzon-readObj-8                1203966               993 ns/op             144 B/op         14 allocs/op
Benchmark_10Fields_Iterator_Object/jzon-readObjCB-8              1000000              1010 ns/op             144 B/op         14 allocs/op
Benchmark_10Fields_Unmarshal_Interface/json-8                     260616              4834 ns/op            1414 B/op         36 allocs/op
Benchmark_10Fields_Unmarshal_Interface/jsoniter-8                 307460              3525 ns/op            1350 B/op         38 allocs/op
Benchmark_10Fields_Unmarshal_Interface/jsoniter-compatible-8      333093              3581 ns/op            1350 B/op         38 allocs/op
Benchmark_10Fields_Unmarshal_Interface/djson-8                    599166              2327 ns/op            1174 B/op         27 allocs/op
Benchmark_10Fields_Unmarshal_Interface/easyjson-8                 428072              2927 ns/op            1174 B/op         27 allocs/op
Benchmark_10Fields_Unmarshal_Interface/ujson-8                    374734              3290 ns/op            1494 B/op         38 allocs/op
Benchmark_10Fields_Unmarshal_Interface/ugorji-8                   244724              4960 ns/op            2221 B/op         36 allocs/op
Benchmark_10Fields_Unmarshal_Interface/jzon-8                     428084              2812 ns/op            1190 B/op         28 allocs/op
Benchmark_10Fields_Unmarshal_Interface/jzon-fast-8                443917              2773 ns/op            1190 B/op         28 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/json-8                 272451              4488 ns/op             432 B/op         14 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/jsoniter-8             999300              1035 ns/op             192 B/op          5 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/jsoniter-compatible-8                 1000000              1011 ns/op             192 B/op          5 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/ugorji-8                               521652              2285 ns/op             832 B/op          7 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/jzon-8                                1000000              1073 ns/op             192 B/op          5 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/jzon-fast-8                           1000000              1088 ns/op             192 B/op          5 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/json-8                                   249818              4807 ns/op             432 B/op         14 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/jsoniter-8                               800250              1484 ns/op             256 B/op         15 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/jsoniter-compatible-8                    856506              1476 ns/op             256 B/op         15 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/jzon-8                                   799386              1529 ns/op             192 B/op          5 allocs/op
Benchmark_20Fields_Unmarshal_Interface/json-8                                     114210              9516 ns/op            3004 B/op         67 allocs/op
Benchmark_20Fields_Unmarshal_Interface/jsoniter-8                                 162044              7528 ns/op            3051 B/op         73 allocs/op
Benchmark_20Fields_Unmarshal_Interface/jsoniter-compatible-8                      157784              7427 ns/op            3051 B/op         73 allocs/op
Benchmark_20Fields_Unmarshal_Interface/djson-8                                    255135              4707 ns/op            2716 B/op         52 allocs/op
Benchmark_20Fields_Unmarshal_Interface/easyjson-8                                 193371              6174 ns/op            2716 B/op         52 allocs/op
Benchmark_20Fields_Unmarshal_Interface/ujson-8                                    171274              7111 ns/op            3340 B/op         73 allocs/op
Benchmark_20Fields_Unmarshal_Interface/ugorji-8                                   119914              9847 ns/op            3764 B/op         61 allocs/op
Benchmark_20Fields_Unmarshal_Interface/jzon-8                                     199820              5925 ns/op            2732 B/op         53 allocs/op
Benchmark_20Fields_Unmarshal_Interface/jzon-fast-8                                206701              5911 ns/op            2732 B/op         53 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/json-8                                 137854              8783 ns/op             648 B/op         24 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/jsoniter-8                             428254              2816 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/jsoniter-compatible-8                  428072              2730 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/ugorji-8                               544728              2298 ns/op             832 B/op          7 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/jzon-8                                 521070              2381 ns/op             368 B/op          9 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/jzon-fast-8                            499662              2425 ns/op             368 B/op          9 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/json-8                                   119913             10048 ns/op             648 B/op         24 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/jsoniter-8                               444157              2758 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/jsoniter-compatible-8                    443912              2852 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/jzon-8                                   244792              4963 ns/op            1075 B/op          9 allocs/op
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
