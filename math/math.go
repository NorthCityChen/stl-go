/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-04-03 09:32:48
 * @Description:
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */
package math

import "math"

type Number interface {
	int | int32 | int64 | float32 | float64
}

// 返回n个数中的最小值
func Max[T Number](arg T, args ...T) T {
	var maxNum T = arg

	for _, val := range args {
		if maxNum < val {
			maxNum = val
		}
	}
	return maxNum
}

// 返回n个数中的最大值
func Min[T Number](arg T, args ...T) T {
	var minNum T = arg
	for _, val := range args {
		if minNum > val {
			minNum = val
		}
	}
	return minNum
}

// 返回n个数的和
func Sum[T Number](args ...T) (sum T) {
	for _, val := range args {
		sum += val
	}
	return sum
}

// 返回num的绝对值
func Abs[T Number](num T) T {
	if num < 0 {
		num *= -1
	}
	return num
}

// 返回 2**n
func Exp2(num int) int64 {
	var ans int64 = 1
	for i := 0; i < num; i++ {
		ans <<= 1
	}
	return ans
}

// 返回x的平方根，精度比原生略低
func Sqrt(x float64) float64 {
	if x < 0 {
		return math.NaN()
	}

	if x == 0 {
		return 0
	}
	last, res := 0.0, 1.0
	for last != res {
		last = res
		res = (res + x/res) / 2
	}
	return res
}
