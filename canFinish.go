package main

import "fmt"

// https://leetcode.cn/problems/course-schedule
func run207() {
	fmt.Println(canFinish(4, [][]int{{1, 2}, {2, 3}, {3, 1}}))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 存储有向图，k: 学习的课程，v: 被下游课程依赖
	prequires := make(map[int][]int, 0)
	// 存储入度, k: 学习的课程，v: 依赖上游可能数目
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
		k := learned[i]
		requires, _ := prequires[k]
		for _, require := range requires {
			indeg[require] -= 1
			if indeg[require] == 0 {
				learned = append(learned, require)
			}
		}
	}

	return len(learned) == numCourses
}
