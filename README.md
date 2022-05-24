# golang_last_stone_weight

You are given an array of integers `stones` where `stones[i]` is the weight of the `ith` stone.

We are playing a game with the stones. On each turn, we choose the **heaviest two stones** and smash them together. Suppose the heaviest two stones have weights `x` and `y` with `x <= y`. The result of this smash is:

- If `x == y`, both stones are destroyed, and
- If `x != y`, the stone of weight `x` is destroyed, and the stone of weight `y` has new weight `y - x`.

At the end of the game, there is **at most one** stone left.

Return *the smallest possible weight of the left stone*. If there are no stones left, return `0`.

## Examples

**Example 1:**

```
Input: stones = [2,7,4,1,8,1]
Output: 1
Explanation:
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.

```

**Example 2:**

```
Input: stones = [1]
Output: 1

```

**Constraints:**

- `1 <= stones.length <= 30`
- `1 <= stones[i] <= 1000`

## 解析

題目給一個整數陣列 stones ，其中 stones[i] 代表第 i 個石頭的重量

每次取出最重的兩個石頭重量分別是 y, x 做相撞

相撞規則如下

如果 x == y ， 則兩顆石頭不存在

如果 y > x , 則把 y-x 重量的石頭放入原本石頭堆中

求剩下一顆以下石頭時，最後的重量是多少

直覺的作法需要把石頭由重到輕排序一次，使用 heap sort 需要 O(nlogn)

每次相撞 2 顆石頭取出最前面 2 顆

然後根據相撞剩下的重量把兩顆放入對應的位置 假設使用 binary search 是 logn

這樣時間複雜度 O(nlogn)

空間複雜度需要額外的 heap O(n)

要降低空間複雜度可以直接把 array 放入 Max-Heap

每次放入 需要 logn

一共n個 所以是 O(nlogn)

每次 pop 出兩個 element, 因為有 heapify 所以是 2 *logN 

然後放入 logN

所以 n 次這樣時間複雜度是 O(nlogn)

空間複雜度因為本身只放入 heap 除了本身的 heap O(n) 並無其他空間複雜

兩個作法都是需要 heap

而第二個作法比較簡化

![](https://i.imgur.com/tnogT2C.png)

## 程式碼
```go
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

```
## 困難點

1. 理解 Max Heap 原理來處理會得到比較好的執行時間

## Solve Point

- [x]  Understand what problem need to solve
- [x]  Analysis Complexity