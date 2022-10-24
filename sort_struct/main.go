package main

import (
	"fmt"
	"sort"
)

type Region struct {
	Id   int
	Name string
	Num  string
}

func main() {
	var rs = []Region{{
		Name: "C",
	}, {
		Name: "A",
	}, {
		Name: "B",
	}}

	sort.Slice(rs, func(i, j int) (less bool) {
		return rs[i].Name < rs[j].Name
	})

	fmt.Printf("%+v\n", rs)
	// Output:
	//   [{ID:0 Name:A Num:} {ID:0 Name:B Num:} {ID:0 Name:C Num:}]
}
