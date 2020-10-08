package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"golang.org/x/tools/benchmark/parse"
)

type resList []*parse.Benchmark

var _ sort.Interface = resList{}

func (l resList) Len() int {
	return len(l)
}

func (l resList) Less(i, j int) bool {
	return l[i].NsPerOp < l[j].NsPerOp
}

func (l resList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

const (
	head  = "|     |     |     | ns/op | B/op | allocs/op |"
	head2 = "| --- | --- | --- | ----- | ---- | --------- |"
	sep   = "|     |     |     |       |      |           |"
)

func parseSet(filename string) parse.Set {
	fp, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error while opening %q: %s", filename, err)
	}
	defer fp.Close()

	set, err := parse.ParseSet(fp)
	if err != nil {
		log.Fatalf("error while parsing: %v", err)
	}
	return set
}

func main() {
	set := parseSet(os.Args[1])

	var keys []string
	m := map[string]resList{}
	for name, arr := range set {
		nameArr := strings.SplitN(strings.TrimPrefix(name, "Benchmark_"), "/", 2)
		if len(nameArr) != 2 {
			log.Fatalf("invalid test name %q", name)
		}
		subName := nameArr[0]

		if len(arr) != 1 {
			log.Fatalf("duplicate test %q", name)
		}
		bench := arr[0]
		bench.Name = nameArr[1]

		if _, ok := m[subName]; !ok {
			keys = append(keys, subName)
		}
		m[subName] = append(m[subName], bench)
	}

	fmt.Println(head)
	fmt.Println(head2)

	sort.Strings(keys)
	for _, k := range keys {
		arr := m[k]
		sort.Sort(arr)
		for _, item := range arr {
			fmt.Printf("| %s | %s | %d | %.0f | %d | %d |\n", k, item.Name,
				item.N, item.NsPerOp, item.AllocedBytesPerOp, item.AllocsPerOp)
		}
		fmt.Println(sep)
	}
}
