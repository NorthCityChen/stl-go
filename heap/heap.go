/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-04-26 16:06:24
 * @Description:
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */
package heap

type Number interface {
	int | int64 | float32 | float64
}

type heap[T Number] struct {
	array []T
	len   int
	cmp   func(T, T) bool
}

func (h *heap[T]) up(x int) {
	for x > 1 && h.cmp(h.array[x], h.array[x/2]) {
		h.array[x], h.array[x/2] = h.array[x/2], h.array[x]
		x /= 2
	}
}

func (h *heap[T]) down(x int) {
	for x*2 <= h.len {
		t := x * 2
		if t+1 <= h.len && h.cmp(h.array[t+1], h.array[t]) {
			t++
		}
		if !h.cmp(h.array[t], h.array[x]) {
			break
		}
		h.array[x], h.array[t] = h.array[t], h.array[x]
		x = t
	}
}

func (h *heap[T]) Push(v T) {
	h.array = append(h.array, v)
	h.len++
	h.up(h.len)
}

func (h *heap[T]) Pop() T {
	if h.len == 0 {
		return 0
	}
	ret := h.array[1]
	h.array[1], h.array[h.len] = h.array[h.len], h.array[1]
	h.array = h.array[:h.len]
	h.len--
	h.down(1)
	return ret
}

func (h *heap[T]) Top() T {
	return h.array[1]
}

func Init[T Number](array []T, cmp func(T, T) bool) heap[T] {
	h := heap[T]{
		array: append([]T{0}, array...),
		len:   len(array),
		cmp:   cmp,
	}
	for i := h.len; i >= 1; i-- {
		h.down(i)
	}
	return h
}
