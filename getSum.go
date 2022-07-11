package main

// https://leetcode.cn/problems/sum-of-two-integers/
func run371() {}

func getSum(a int, b int) int {
	carry := int32(0)
	c := int32(0)
	for i := 0; i < 32; i++ {
		// 第i为，如果a， b， carry出现0个或者2个1， 则为1，出现1个或者3个1，则为0
		bit := int32(((a >> i) ^ (b >> i)) & 1)
		bit ^= carry
		// // 第i为，如果a， b， carry出现两个以上的1，则进位了
		if carry == 1 {
			if ((a>>i)&1) == 1 || ((b>>i)&1) == 1 {
				carry = 1
			} else {
				carry = 0
			}
		} else {
			if ((a>>i)&1) == 1 && ((b>>i)&1) == 1 {
				carry = 1
			} else {
				carry = 0
			}
		}
		c |= (bit << i)
	}
	return int(c)
}
