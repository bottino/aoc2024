package day16

import (
	"container/heap"
)

type Item[T comparable] struct {
	value T
	rank  int
	index int
}

// Implements heap.Interface
type PriorityQueue[T comparable] struct {
	q      []*Item[T]
	lookup map[T]*Item[T]
}

func (pq *PriorityQueue[T]) Len() int {
	return len((*pq).q)
}

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	return (*pq).q[i].rank < (*pq).q[j].rank
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	(*pq).q[i], (*pq).q[j] = (*pq).q[j], (*pq).q[i]
	(*pq).q[i].index = i
	(*pq).q[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	item := x.(*Item[T])
	item.index = (*pq).Len()
	(*pq).lookup[item.value] = item
	(*pq).q = append((*pq).q, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	oldq := (*pq).q
	n := (*pq).Len()
	item := (*pq).q[n-1]
	(*pq).q = oldq[:n-1]
	delete((*pq).lookup, item.value)
	return item
}

// Adds an item if doesn't exist, updates the rank if it does
func (pq *PriorityQueue[T]) AddWithRank(value T, rank int) {
	if pItem, ok := (*pq).lookup[value]; ok {
		pItem.rank = rank
		heap.Fix(pq, pItem.index)
	} else {
		item := Item[T]{value: value, rank: rank}
		heap.Push(pq, &item)
	}
}

func (pq *PriorityQueue[T]) PopMin() T {
	var popped *Item[T]
	popped = heap.Pop(pq).(*Item[T])
	return popped.value
}

func NewPQueue[T comparable]() PriorityQueue[T] {
	pq := PriorityQueue[T]{
		make([]*Item[T], 0, 0),
		make(map[T]*Item[T]),
	}
	heap.Init(&pq)
	return pq
}
