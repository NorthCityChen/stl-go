/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-04-01 21:46:15
 * @Description:
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */
package queue

type Number interface {
	int | int64 | float32 | float64
}

type Queue[T Number] struct {
	size  int
	front *QueueNode[T]
	end   *QueueNode[T]
}

type QueueNode[T Number] struct {
	Val  T
	Next *QueueNode[T]
}

func Init[T Number]() Queue[T] {
	return Queue[T]{size: 0, front: nil, end: nil}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) Push(val T) bool {
	newNode := QueueNode[T]{Val: val, Next: nil}
	q.size++
	q.end.Next = &newNode
	q.end = &newNode
	return true
}

func (q *Queue[T]) GetFront() T { return q.front.Val }

func (q *Queue[T]) Pop() (T, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	retVal := q.front.Val
	q.front = q.front.Next
	return retVal, true
}
