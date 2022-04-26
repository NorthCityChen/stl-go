/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-04-26 17:08:18
 * @Description: Binary Indexed Trees
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */

package bit

type Number interface {
	int | int64 | float32 | float64
}

type bit[T Number] struct {
	array []T
	brray []T
	len   int
}

func lowbit(x int) int {
	return x & -x
}

// 初始化一个树状数组
func Init[T Number](array []T) bit[T] {
	ret := bit[T]{
		array: append([]T{0}, array...),
		brray: make([]T, len(array)+5),
		len:   len(array),
	}

	for i := 1; i < len(ret.array); i++ {
		ret.Add(i, ret.array[i])
	}
	return ret
}

// 在第x个数的位置+k
func (b *bit[T]) Add(x int, k T) {
	for x <= b.len {
		b.brray[x] = b.brray[x] + k
		x = x + lowbit(x) // 父节点
	}
}

// 求第一个数到第x个数的和
func (b *bit[T]) Ask(x int) T {
	var ans T = 0
	for x >= 1 {
		ans = ans + b.brray[x]
		x = x - lowbit(x)
	}
	return ans
}
