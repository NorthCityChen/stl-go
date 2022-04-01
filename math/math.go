/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-04-01 21:04:53
 * @Description:
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */
package math

import "math"

type Number interface {
	int | int32 | int64 | float32 | float64
}

func Inf[T Number]() T {
	return T(math.Inf(1))
}

func Max[T Number](arg T, args ...T) T {
	var maxNum T = arg

	for _, val := range args {
		if maxNum < val {
			maxNum = val
		}
	}
	return maxNum
}

func Min[T Number](arg T, args ...T) T {
	var minNum T = arg
	for _, val := range args {
		if minNum > val {
			minNum = val
		}
	}
	return minNum
}

func Sum[T Number](args ...T) (sum T) {
	for _, val := range args {
		sum += val
	}
	return sum
}

func Abs[T Number](num T) T {
	if num < 0 {
		num *= -1
	}
	return num
}

// sort
