/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-05-19 22:40:55
 * @Description:
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */
package stack

import "github.com/NorthCityChen/stl-go/dequeue"

type Number interface {
	int | int64 | float32 | float64
}

type stack[T Number] struct {
	content *dequeue.Dequeue[T]
}

func Init[T Number]() *stack[T] {
	return &stack[T]{content: dequeue.Init[T]()}
}

func (s *stack[T]) IsEmpty() bool {
	return s.content.IsEmpty()
}

func (s *stack[T]) Size() int {
	return s.content.Size()
}

func (s *stack[T]) Push(val T) bool {
	return s.content.LPush(val)
}

func (s *stack[T]) Pop() (T, bool) {
	return s.content.LPop()
}

func (s *stack[T]) Top() (T, bool) {
	return s.content.LNode()
}

func (s *stack[T]) Clear() bool {
	return s.content.Clear()
}
