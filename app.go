/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-04-26 17:05:06
 * @Description:
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */
package main

import (
	"fmt"

	"gitee.com/NorthCityChen/stl-go/bit"
)

var a = []int{1, 2, 65, 4, 33, 5}

func main() {
	ans := a[0]
	for i := len(a) - 1; i >= 1; i-- {
		a[i] = a[i] - a[i-1]
	}
	b := bit.Init(a)
	fmt.Println(b)
	ans = b.Ask(6)
	fmt.Println(ans)
}
