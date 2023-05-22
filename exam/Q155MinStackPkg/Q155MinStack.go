package Q155MinStackPkg

import "math"

type MinStack struct {
	s []int
}

func Constructor() MinStack {
	return MinStack{
		s: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.s = append(this.s, val)
}

func (this *MinStack) Pop() {
	this.s = this.s[:len(this.s)-1]
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	min := math.MaxInt32
	for _, v := range this.s {
		if v < min {
			min = v
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
