/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-04-26 14:25:09
 * @Description:
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */
package main

import (
	"fmt"

	"gitee.com/NorthCityChen/stl-go/heap"
)

var a = []int{1, 2, 65, 4, 5}

func main() {
	h := heap.Init(a, false)
	fmt.Println(h)
	h.Push(77)
	fmt.Println(h)
	fmt.Println(h.Top())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
}
