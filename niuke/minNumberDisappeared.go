package niuke

// https://www.nowcoder.com/practice/50ec6a5b0e4e45348544348278cdcee5?tpId=295&tqId=2188893&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
func Run_BMP53() {}

/*
思路：

nnn个长度的数组，没有重复，则如果数组填满了1～n1～n1～n，那么缺失n+1n+1n+1，如果数组填不满1～n1～n1～n，那么缺失的就是1～n1～n1～n中的数字。对于这种快速查询某个元素是否出现过的问题，还是可以使用哈希表快速判断某个数字是否出现过。

具体做法：

step 1：构建一个哈希表，用于记录数组中出现的数字。
step 2：从1开始，遍历到n，查询哈希表中是否有这个数字，如果没有，说明它就是数组缺失的第一个正整数，即找到。
step 3：如果遍历到最后都在哈希表中出现过了，那缺失的就是n+1.
*/

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums int整型一维数组
 * @return int整型
 */
func minNumberDisappeared(nums []int) int {
	// write code here
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1 //用n+1表示是因为n要指向nums[n-1]
		}
	}

	for i := 0; i < n; i++ {
		k := nums[i]
		if k < 0 {
			k = -k
		}
		k-- //下标0~(n-1)代表数字1~n
		if k < n && nums[k] > 0 {
			nums[k] = -nums[k]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}
