package Q703KthLargestElementinaStreamPkg

import "container/heap"

type KthLargest struct {
	minHeap *IntHeap
	k       int
}

func Constructor(k int, nums []int) KthLargest {
	h := append(IntHeap{}, nums...)
	heap.Init(&h)

	for len(h) > k {
		heap.Pop(&h)
	}

	obj := KthLargest{
		minHeap: &h,
		k:       k,
	}
	return obj
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.minHeap, val)
	if len(*this.minHeap) > this.k {
		heap.Pop(this.minHeap)
	}
	// fmt.Println(this.minHeap, val, x)

	if len(*this.minHeap) > 0 {
		return (*this.minHeap)[0]
	}

	return 0
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	tmp := old[len(*h)-1]
	*h = old[0 : len(*h)-1]
	return tmp
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
