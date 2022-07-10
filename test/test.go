package test

import (
	"fmt"
	"unsafe"
)

func Test() {
	// 测试slice cut
	i := []int{9, 8, 7, 6, 5}
	fmt.Println(i[:2])
	fmt.Println(i[2:5])

	// 测试基本类型大小
	a := ^uint(0)
	fmt.Println(a)
	max := int(^uint32(0) >> 1)
	min := ^max
	fmt.Println(min)
	fmt.Println(max)
	fmt.Println(unsafe.Sizeof(max))
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(int64(16)))
	fmt.Println(unsafe.Sizeof(int32(16)))
	fmt.Println(int32(max + 1))
	fmt.Println(int(max + 1))

	// 测试数组内存分配
	num := []int{1, 2, 3}
	test1(num)
	test2(num)
	fmt.Println(num)

	// test 切片内存分配
	nums := make([]int, 0)
	nums = append(nums, 0)
	nums = append(nums, 1)
	fmt.Println(unsafe.Pointer(&nums))
	test3(nums)

	// test数组内存分配
	var eles [2]int
	eles[0] = 0
	eles[1] = 1
	fmt.Println(unsafe.Pointer(&eles))
	test4(eles)
	fmt.Println(eles[0])

	// test切片扩容
	nums = make([]int, 0, 2)
	nums = append(nums, 0)
	nums = append(nums, 1)
	fmt.Println(unsafe.Pointer(&nums))
	test5(nums)
	fmt.Println(nums[0])
	// 指针未变化，指向内容仍是原来的nums，函数内的指针修改指向新地址，不影响这个指针
	fmt.Println(unsafe.Pointer(&nums))

	fmt.Println("~~~~~")
	mp := make(map[int]int)
	mp[1] = 1
	fmt.Println(unsafe.Pointer(&mp))
	mp[2] = 2
	fmt.Println(unsafe.Pointer(&mp))
	testMap(mp)
	fmt.Println(unsafe.Pointer(&mp))
	fmt.Println(mp[1])

	// 8进制和16进制
	n8 := 017
	n16 := 0x17A
	fmt.Println(n8)
	fmt.Println(n16)

	// 测试字符串入参
	str := []byte("abcdef")
	testStrReOrder(str)
	fmt.Println(str)
}

func testMap(mp map[int]int) {
	fmt.Println(unsafe.Pointer(&mp))
	mp[3] = 3
	mp[1] = 4
}
func testStrReOrder(str []byte) {
	str[0], str[1] = str[1], str[0]
	str = append(str, 'g')
}

func test5(nums []int) {
	fmt.Println(unsafe.Pointer(&nums))
	// 容量增加，导致nums的实际存储地址变化，但&nums指针未变，入参的原指针不指向新地址了
	nums = append(nums, 3)
	nums[0] = 5
	fmt.Println(unsafe.Pointer(&nums))
}

func test3(nums []int) {
	fmt.Println(unsafe.Pointer(&nums))
}

func test4(eles [2]int) {
	fmt.Println(unsafe.Pointer(&eles))
	eles[0] = 5
}

func test1(num []int) {
	num[2] = 9
}

func test2(num []int) {
	num[1] = 5
}
