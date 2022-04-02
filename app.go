/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-04-02 12:42:40
 * @Description:
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */
package main

import (
	"fmt"
	"stl-go/dequeue"
)

func main() {
	q := dequeue.Init[int]()
	q.LPush(12)
	q.LPush(23)
	q.LPush(34)
	q.Clear()
	q.LPush(34)
	q.RPush(44)
	fmt.Println(q.LPop())
	fmt.Println(q.LPop())
	fmt.Println(q.LPop())
	fmt.Println(q.LPop())
	// q := queue.Init[int]()
	// q.Push(12)
	// q.Push(23)
	// q.Push(44)
	// // q.IsEmpty()
	// fmt.Println(q.Pop())
	// fmt.Println(q.Pop())
	// fmt.Println(q.Pop())
	// fmt.Println(q.Pop())
}
