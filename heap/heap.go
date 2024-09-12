package main

import (
	"container/heap"
	"fmt"
)

// 定义一个包含多个字段的结构体
type Item struct {
	value    string // 元素的值
	priority int    // 元素的优先级（优先级越低，越靠近堆顶）
}

// 定义一个基于 Item 的最小堆
type PriorityQueue []*Item

// 实现 heap.Interface 接口
func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 优先级较小的元素优先级较高
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	// 向队列中插入元素
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	// 移除并返回队列中优先级最高的元素
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	// 创建优先队列并初始化
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// 向优先队列中添加元素
	heap.Push(&pq, &Item{value: "task1", priority: 3})
	heap.Push(&pq, &Item{value: "task2", priority: 1})
	heap.Push(&pq, &Item{value: "task3", priority: 2})

	// 按优先级顺序取出元素
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("Value: %s, Priority: %d\n", item.value, item.priority)
	}
}

