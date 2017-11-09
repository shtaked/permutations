package priorityqueue

import "container/heap"

// An item is something we manage in a priority queue.
type item struct {
	value    interface{} // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A queue implements heap.Interface and holds Items.
type queue []*item

func (pq queue) Len() int { return len(pq) }

func (pq queue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq queue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *queue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *queue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type Queue struct {
	q queue
}

func NewQueue() *Queue{
	q := make(queue, 0)
	heap.Init(&q)
	return &Queue{q}
}

func (pq *Queue) Push(value interface{}, priority int) {
	heap.Push(&pq.q, &item{value:value, priority:priority, index:len(pq.q)})
}

func (pq *Queue) Pop() (value interface{}, priority int) {
	p := heap.Pop(&pq.q).(*item)
	return p.value, p.priority
}

func (pq *Queue) Len() int {
	return pq.q.Len()
}