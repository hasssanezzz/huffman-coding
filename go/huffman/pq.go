package huffman

type PriorityQueue []*Node

// Less compares items based on their priorities.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Freq < pq[j].Freq
}

// Len returns the number of items in the priority queue.
func (pq PriorityQueue) Len() int { return len(pq) }

// Push adds an item to the priority queue.
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	item.Index = len(*pq)
	*pq = append(*pq, item)
}

// Swap swaps two items in the priority queue.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Pop removes and returns the item with the highest priority.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
