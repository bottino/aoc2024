package day16

import "container/heap"

type Item[T comparable] struct {
	value T
	rank  int
	index int
}

// Implements heap.Interface
type PriorityQueue[T comparable] []*Item[T]

func (pq *PriorityQueue[T]) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	return (*pq)[i].rank < (*pq)[j].rank
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	item := x.(*Item[T])
	item.index = (*pq).Len()
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := (*pq).Len()
	item := (*pq)[n-1]
	*pq = old[:n-1]
	return item
}

func (pq *PriorityQueue[T]) AddWithRank(value T, rank int) {
	item := Item[T]{value: value, rank: rank}
	heap.Push(pq, &item)
}

func (pq *PriorityQueue[T]) PopMin() T {
	var popped *Item[T]
	popped = heap.Pop(pq).(*Item[T])
	return popped.value
}

func NewPQueue[T comparable]() PriorityQueue[T] {
	pq := make(PriorityQueue[T], 0, 0)
	heap.Init(&pq)
	return pq
}
