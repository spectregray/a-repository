package utils

type MinHeap []int

func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

// function Len() for type MinHeap with alias mh that returns int
// what's the point of this?
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