/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-05-03 18:51:46
 * @Description:
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */
package main

import (
	"fmt"

	"gitee.com/NorthCityChen/stl-go/math"
)

var a = []int{1, 2, 65, 4, 33, 5}

func main() {
	// h := heap.Init(a, heap.MinHeap[int]())
	// fmt.Println(h)

	ans := math.C(8, 2)
	fmt.Println(ans)

	ans = math.A(12, 2)
	fmt.Println(ans)
}
