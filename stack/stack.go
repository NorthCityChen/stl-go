/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-04-01 21:35:59
 * @Description:
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */
package stack

type Number interface {
	int | int64 | float32 | float64
}

type Stack[T Number] struct {
	size T
	top  *stackItem[T]
}

type stackItem[T Number] struct {
	Val  T
	Next *stackItem[T]
}

func Init[T Number]() Stack[T] {
	return Stack[T]{size: 0, top: nil}
}

func (s *Stack[T]) IsEmpty() bool { return s.size == 0 }

func (s *Stack[T]) Push(val T) bool {
	newNode := stackItem[T]{Val: val, Next: s.top}
	s.size++
	s.top = &newNode
	return true
}

func (s *Stack[T]) GetTop() T {
	return s.top.Val
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	retVal := s.top.Val
	s.size--
	s.top = s.top.Next
	return retVal, true
}
