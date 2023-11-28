package utils

type MinHeap []int

// Interface stuffs
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h MinHeap) Len() int {
	return len(h)
}

func (h *MinHeap) Push(x interface{}) { // interace{} is the any type
	*h = append(*h, x.(int)) // x.(int) is type assertion
}

func (h *MinHeap) Pop() interface{} {
	oldHeap := *h 
	last := oldHeap[len(oldHeap) - 1]
	*h = oldHeap[:len(oldHeap) - 1]
	return last
}

// Non interface stuffs
func (h *MinHeap) Peek() interface{} {
	heap := *h 
	if len(heap) > 0 {
		return heap[0]
	}
	return 0
}