/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-04-17 12:04:32
 * @Description:
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */
package dequeue

type Number interface {
	int | int64 | float32 | float64 | string
}

// dequeue: Based on Double Linked List
type Dequeue[T Number] struct {
	size      int
	leftHead  *dequeueNode[T]
	rightHead *dequeueNode[T]
}

type dequeueNode[T Number] struct {
	Val      T
	LPointer *dequeueNode[T]
	RPointer *dequeueNode[T]
}

func Init[T Number]() *Dequeue[T] {
	return &Dequeue[T]{size: 0, leftHead: new(dequeueNode[T]), rightHead: new(dequeueNode[T])}
}

func (q *Dequeue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Dequeue[T]) Size() int {
	return q.size
}

func (q *Dequeue[T]) LPush(val T) bool {
	newNode := &dequeueNode[T]{Val: val, LPointer: new(dequeueNode[T]), RPointer: q.leftHead}
	if q.IsEmpty() {
		q.leftHead = newNode
		q.rightHead = newNode
	}
	q.size++
	q.leftHead.LPointer = newNode
	q.leftHead = newNode
	return true
}

func (q *Dequeue[T]) LPop() (T, bool) {
	if q.IsEmpty() {
		return T(rune(0)), false
	}
	val := q.leftHead.Val
	q.leftHead = q.leftHead.RPointer
	q.size--
	return val, true
}

func (q *Dequeue[T]) RPush(val T) bool {
	newNode := &dequeueNode[T]{Val: val, LPointer: q.rightHead, RPointer: new(dequeueNode[T])}
	if q.IsEmpty() {
		q.leftHead = newNode
		q.rightHead = newNode
	}
	q.size++
	q.rightHead.RPointer = newNode
	q.rightHead = newNode
	return true
}

func (q *Dequeue[T]) RPop() (T, bool) {
	if q.IsEmpty() {
		return T(rune(0)), false
	}
	val := q.rightHead.Val
	q.rightHead = q.rightHead.LPointer
	q.size--
	return val, true
}

func (q *Dequeue[T]) LNode() (T, bool) {
	if q.IsEmpty() {
		return T(rune(0)), false
	}
	return q.leftHead.Val, true
}

func (q *Dequeue[T]) RNode() (T, bool) {
	if q.IsEmpty() {
		return T(rune(0)), false
	}
	return q.rightHead.Val, true
}

func (q *Dequeue[T]) Clear() bool {
	for !q.IsEmpty() {
		q.RPop()
	}
	return q.IsEmpty()
}
