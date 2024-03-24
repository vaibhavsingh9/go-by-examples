package main

type MinHeap []int

func (h *MinHeap) BuildHeap(array []int) {
	lastNonLeafNodeIdx := (len(array) - 2) / 2
	for currentIdx := lastNonLeafNodeIdx; currentIdx >= 0; currentIdx-- {
		endIdx := len(array) - 1
		h.shiftDown(currentIdx, endIdx)
	}
}

func (h *MinHeap) shiftDown(currentIdx, endIdx int) {
	leftChildIdx := currentIdx*2 + 1
	leftChildIdx <= endIdx{
		rightChildIdx :=
	}
}

func main() {
	array := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}
	minHeap := &MinHeap{}

	*minHeap = array
	minHeap.BuildHeap(array)
}
