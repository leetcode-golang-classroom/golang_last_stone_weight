package sol

import "container/heap"

type IntMaxHeap []int

// Len() int
func (h *IntMaxHeap) Len() int {
	return len(*h)
}

// Less(i, j int) bool
func (h *IntMaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

// Swap(i, j int)
func (h *IntMaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntMaxHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func lastStoneWeight(stones []int) int {
	priorityQueue := &IntMaxHeap{}
	sLen := len(stones)
	heap.Init(priorityQueue)
	for idx := 0; idx < sLen; idx++ {
		heap.Push(priorityQueue, stones[idx])
	}
	result := 0
	for priorityQueue.Len() >= 2 {
		first := heap.Pop(priorityQueue).(int)
		second := heap.Pop(priorityQueue).(int)
		val := first - second
		if val > 0 {
			heap.Push(priorityQueue, val)
		}
	}
	if priorityQueue.Len() > 0 {
		result = heap.Pop(priorityQueue).(int)
	}
	return result
}
