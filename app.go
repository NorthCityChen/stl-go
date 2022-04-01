/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-04-01 21:47:35
 * @Description:
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */
package main

import (
	"stl-go/queue"
)

func main() {
	// s := stack.Init[int]()
	// s.Push(12)
	// s.Push(13)
	// fmt.Println(s.GetTop())
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	q := queue.Init[int]()
	q.Push(12)
	q.Push(34)
	// fmt.Println(q.Pop())
	// fmt.Println(q.Pop())
}
