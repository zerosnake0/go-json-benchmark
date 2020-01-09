package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

var s = `
Benchmark_10Fields_Iterator_Object/jsoniter-readObj-8            1243933              1072 ns/op             144 B/op         14 allocs/op
Benchmark_10Fields_Iterator_Object/jsoniter-readObjCB-8          1000000              1126 ns/op             144 B/op         14 allocs/op
Benchmark_10Fields_Iterator_Object/jsonparser-8                   856482              1171 ns/op              80 B/op          4 allocs/op
Benchmark_10Fields_Iterator_Object/jzon-readObj-8                 749488              1386 ns/op             144 B/op         14 allocs/op
Benchmark_10Fields_Iterator_Object/jzon-readObjCB-8               856426              1320 ns/op             144 B/op         14 allocs/op
Benchmark_10Fields_Unmarshal_Interface/json-8                     235119              5346 ns/op            1414 B/op         36 allocs/op
Benchmark_10Fields_Unmarshal_Interface/jsoniter-8                 333102              4008 ns/op            1350 B/op         38 allocs/op
Benchmark_10Fields_Unmarshal_Interface/jsoniter-compatible-8      333104              3908 ns/op            1350 B/op         38 allocs/op
Benchmark_10Fields_Unmarshal_Interface/djson-8                    521389              2464 ns/op            1174 B/op         27 allocs/op
Benchmark_10Fields_Unmarshal_Interface/easyjson-8                 428254              2979 ns/op            1174 B/op         27 allocs/op
Benchmark_10Fields_Unmarshal_Interface/ujson-8                    333092              3515 ns/op            1494 B/op         38 allocs/op
Benchmark_10Fields_Unmarshal_Interface/ugorji-8                   223449              5316 ns/op            2222 B/op         36 allocs/op
Benchmark_10Fields_Unmarshal_Interface/jzon-8                     333105              3266 ns/op            1190 B/op         28 allocs/op
Benchmark_10Fields_Unmarshal_Interface/jzon-fast-8                374584              3080 ns/op            1190 B/op         28 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/json-8                 249820              4743 ns/op             432 B/op         14 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/jsoniter-8            1000000              1125 ns/op             192 B/op          5 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/jsoniter-compatible-8                 1000000              1125 ns/op             192 B/op          5 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/ugorji-8                               444151              2584 ns/op             832 B/op          7 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/jzon-8                                 749409              1679 ns/op             212 B/op         15 allocs/op
Benchmark_10Fields_Unmarshal_StructWithTag/jzon-fast-8                            750032              1757 ns/op             212 B/op         15 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/json-8                                   235126              5082 ns/op             432 B/op         14 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/jsoniter-8                               705384              1539 ns/op             256 B/op         15 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/jsoniter-compatible-8                    749451              1605 ns/op             256 B/op         15 allocs/op
Benchmark_10Fields_Unmarshal_StructWoTag/jzon-8                                   571058              1978 ns/op             212 B/op         15 allocs/op
Benchmark_20Fields_Unmarshal_Interface/json-8                                     115303             10319 ns/op            3003 B/op         67 allocs/op
Benchmark_20Fields_Unmarshal_Interface/jsoniter-8                                 157783              7782 ns/op            3052 B/op         73 allocs/op
Benchmark_20Fields_Unmarshal_Interface/jsoniter-compatible-8                      151790              8056 ns/op            3051 B/op         73 allocs/op
Benchmark_20Fields_Unmarshal_Interface/djson-8                                    213811              4961 ns/op            2715 B/op         52 allocs/op
Benchmark_20Fields_Unmarshal_Interface/easyjson-8                                 187364              6815 ns/op            2715 B/op         52 allocs/op
Benchmark_20Fields_Unmarshal_Interface/ujson-8                                    159886              7523 ns/op            3339 B/op         73 allocs/op
Benchmark_20Fields_Unmarshal_Interface/ugorji-8                                   115303             10423 ns/op            3763 B/op         61 allocs/op
Benchmark_20Fields_Unmarshal_Interface/jzon-8                                     181689              6973 ns/op            2732 B/op         53 allocs/op
Benchmark_20Fields_Unmarshal_Interface/jzon-fast-8                                181683              6637 ns/op            2732 B/op         53 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/json-8                                 131776              9926 ns/op             648 B/op         24 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/jsoniter-8                             399579              2973 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/jsoniter-compatible-8                  413475              3016 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/ugorji-8                               499627              2381 ns/op             832 B/op          7 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/jzon-8                                 352695              3235 ns/op             408 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWithTag/jzon-fast-8                            370143              3163 ns/op             408 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/json-8                                   111020             11700 ns/op             648 B/op         24 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/jsoniter-8                               413481              2897 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/jsoniter-compatible-8                    413498              2931 ns/op             512 B/op         29 allocs/op
Benchmark_20Fields_Unmarshal_StructWoTag/jzon-8                                   244722              4678 ns/op             408 B/op         29 allocs/op
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
