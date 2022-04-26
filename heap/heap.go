/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-04-26 14:24:36
 * @Description:
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */
package heap

type Number interface {
	int | int64 | float32 | float64
}

type Heap[T Number] struct {
	array []T
	len   int
	isMax bool
}

func (h *Heap[T]) up(x int) {
	if h.isMax {
		for x > 1 && h.array[x] > h.array[x/2] {
			h.array[x], h.array[x/2] = h.array[x/2], h.array[x]
			x /= 2
		}
	} else {
		for x > 1 && h.array[x] < h.array[x/2] {
			h.array[x], h.array[x/2] = h.array[x/2], h.array[x]
			x /= 2
		}
	}

}

func (h *Heap[T]) down(x int) {
	// fmt.Println(h.len)
	if h.isMax {
		for x*2 <= h.len {
			t := x * 2
			if t+1 <= h.len && h.array[t+1] > h.array[t] {
				t++
			}
			if h.array[t] <= h.array[x] {
				break
			}
			h.array[x], h.array[t] = h.array[t], h.array[x]
			x = t
		}
	} else {
		for x*2 <= h.len {
			t := x * 2
			if t+1 <= h.len && h.array[t+1] > h.array[t] {
				t++
			}
			if h.array[t] >= h.array[x] {
				break
			}
			h.array[x], h.array[t] = h.array[t], h.array[x]
			x = t
		}
	}

}

func (h *Heap[T]) Push(v T) {
	h.array = append(h.array, v)
	h.len++
	h.up(h.len)
}

func (h *Heap[T]) Pop() T {
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

func (h *Heap[T]) Top() T {
	return h.array[1]
}

func Init[T Number](array []T, isMax bool) Heap[T] {
	h := Heap[T]{
		array: append([]T{0}, array...),
		len:   len(array),
		isMax: isMax,
	}
	for i := len(h.array) - 1; i >= 1; i-- {
		h.down(i)
	}
	return h
}
