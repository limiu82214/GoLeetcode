package Q705DesignHashSetPkg

type MyHashSet struct {
	data []int
}

func Constructor() MyHashSet {
	return MyHashSet{
		data: make([]int, 1000001),
	}
}

func (this *MyHashSet) Add(key int) {
	this.data[key] = 1
}

func (this *MyHashSet) Remove(key int) {
	this.data[key] = 0
}

func (this *MyHashSet) Contains(key int) bool {
	return this.data[key] == 1
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
