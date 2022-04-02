/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-04-02 12:40:51
 * @Description:
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */
package queue

import "stl-go/dequeue"

type Number interface {
	int | int64 | float32 | float64
}

// queue: Based on dequeue
type queue[T Number] struct {
	content *dequeue.Dequeue[T]
}

func Init[T Number]() *queue[T] {
	return &queue[T]{content: dequeue.Init[T]()}
}

func (q *queue[T]) Push(val T) bool {
	return q.content.RPush(val)
}

func (q *queue[T]) Pop() (T, bool) {
	return q.content.LPop()
}

func (q *queue[T]) IsEmpty() bool {
	return q.content.IsEmpty()
}

func (q *queue[T]) Size() int {
	return q.Size()
}

func (q *queue[T]) Clear() bool {
	return q.content.Clear()
}

func (q *queue[T]) Front() (T, bool) {
	return q.content.LNode()
}
