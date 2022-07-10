package main

import "fmt"

//https://leetcode.cn/problems/course-schedule-ii/
func run210() {
	// fmt.Println(findOrder(4, [][]int{{1, 2}, {2, 3}}))
	fmt.Println(findOrder(2, [][]int{{0, 1}}))
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	prequires := make(map[int][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, p := range prerequisites {
		prequires[p[1]] = append(prequires[p[1]], p[0])
		indeg[p[0]]++
	}

	learned := make([]int, 0)
	for course, upcourse := range indeg {
		if upcourse == 0 {
			learned = append(learned, course)
		}
	}

	for i := 0; i < len(learned); i++ {
		requires, _ := prequires[learned[i]]
		for _, r := range requires {
			indeg[r]--
			if indeg[r] == 0 {
				learned = append(learned, r)
			}
		}
	}

	if len(learned) != numCourses {
		return nil
	}
	return learned
}
