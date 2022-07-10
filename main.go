package main

import (
	"fmt"
	"sort"
)

func main() {
	// test.Test()
	// jianzhioffer.Run56_1()
	// run110()
	// niuke.Run_BM96()
	// run76()
	strs := []string{"abc", "ab", "bca", "cba", "acb", "bac"}
	sort.Strings(strs)
	fmt.Println(strs)
}
