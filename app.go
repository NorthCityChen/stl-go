/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-04-26 15:55:37
 * @Description:
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */
package main

import (
	"fmt"

	"gitee.com/NorthCityChen/stl-go/heap"
)

// import "container/heap"

var a = []int{1, 2, 65, 4, 5}

func main() {
	// <: 小根堆 >: 大根堆
	f := func(t1, t2 int) bool { return t1 > t2 }
	h := heap.Init(a, f)
	fmt.Println(h)

	fmt.Println(h)

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
}
