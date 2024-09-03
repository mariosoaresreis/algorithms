package MergeIntervals

import "container/heap"

type Job struct {
	Start   int
	End     int
	CPULoad int
}

type PriorityQueue []*Job

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].End < pq[j].End
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Push(x interface{}) {
	pq = append(pq, x.(*Job))
}

func (pq PriorityQueue) Pop() interface{} {
	old := pq
	n := len(old)
	item := old[n-1]
	pq = old[0 : n-1]
	return item
}

func test() {
	pq := PriorityQueue{}
	heap.Init(pq)
}
