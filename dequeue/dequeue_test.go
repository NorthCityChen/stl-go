/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-04-27 13:36:25
 * @Description:
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */
package dequeue

import (
	"testing"
	// "gitee.com/NorthCityChen/stl-go/dequeue"
)

// type Number interface {
// 	int | int64 | float32 | float64 | string
// }

// func zeroValue[T Number]() (value T) { return }
func TRune[T Number]() T { return T(rune(0)) }

func BenchmarkZeroValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zeroValue[float64]()
		zeroValue[float32]()
		zeroValue[int]()
		zeroValue[int64]()
		zeroValue[string]()
	}
}

func BenchmarkTRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TRune[float64]()
		TRune[float32]()
		TRune[int64]()
		TRune[int]()
		TRune[string]()
	}
}

func BenchmarkRunDequeue(b *testing.B) {
	dq := Init[int]()
	for i := 0; i < b.N; i++ {
		dq.LPush(1)
		dq.RPush(2)
		dq.LPop()
		dq.RPop()
	}
}
