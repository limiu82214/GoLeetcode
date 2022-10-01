package main

import (
	"encoding/json"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func Decorate(inter func(), times int, cpurate int) {
	if cpurate == 0 {
		cpurate = 100
	}
	// runtime.SetCPUProfileRate(cpurate)
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("create file error", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("errrrrrrrrrrrrrrrr", err)
	}
	defer pprof.StopCPUProfile()
	runtime.GC() // 決定是否先回收其餘的記憶體使用
	for i := 0; i < times; i++ {
		inter()
	}
	f1, _ := os.Create("mem.prof")
	pprof.WriteHeapProfile(f1)
	defer f1.Close()
}

func JsonStringToSliceAny(s string) []interface{} {
	var slice []interface{}
	json.Unmarshal([]byte(s), &slice)
	return slice
}

func JsonToSlice[E interface{}](s string) []E {
	var slice []E
	json.Unmarshal([]byte(s), &slice)
	return slice
}
func JsonToSliceSlice[E interface{}](s string) [][]E {
	var slice [][]E
	json.Unmarshal([]byte(s), &slice)
	return slice
}

func ConvertSlice[E any](in []any) (out []E) {
	out = make([]E, 0, len(in))
	for _, v := range in {
		out = append(out, v.(E))
	}
	return out
}

/*
func ConvertFloatToInt[A interface{}, B interface{}](in []any) (out []B) {
	out = make([]B, 0, len(in))
	for _, v := range in {
		v2, _ := v.(A)
		fmt.Println(int(v2))
		// v.(B)
		// v.(type) = A.(type)
		// fmt.Println(int(nl.(A.(type))))
		// out = append(out, v.(B))
	}
	return out
}
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Accept complete tree JSONArrayToTreeNode [5,3,6,2,4,null,7] use it
func JSONArrayToTreeNode(j string) *TreeNode {
	arg := JsonStringToSliceAny(j)
	if len(arg) < 1 {
		return nil
	}
	list := make([]*TreeNode, 0)

	for _, v := range arg {
		if v == nil {
			list = append(list, nil)
			continue
		}
		tNode := &TreeNode{
			Val:   int(v.(float64)),
			Left:  nil,
			Right: nil,
		}
		list = append(list, tNode)
	}

	for idx := range list {
		l := len(list)
		if idx*2+1 < l {
			list[idx].Left = list[idx*2+1]
		}
		if idx*2+2 < l {
			list[idx].Right = list[idx*2+2]
		}
	}

	return list[0]
}

// if left is nil skip right node
func JSONArrayToTreeNodeV2(j string) *TreeNode {
	arg := JsonStringToSliceAny(j)
	if len(arg) < 1 {
		return nil
	}
	list := make([]*TreeNode, 0)

	for _, v := range arg {
		if v == nil {
			list = append(list, nil)
			continue
		}
		tNode := &TreeNode{
			Val:   int(v.(float64)),
			Left:  nil,
			Right: nil,
		}
		list = append(list, tNode)
	}
	popList := func() *TreeNode {
		if len(list) == 0 {
			return nil
		}
		node := list[0]
		list = list[1:]
		return node
	}

	head := list[0]
	list = list[1:]
	queue := []*TreeNode{}
	queue = append(queue, head)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		node.Left = popList()
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		node.Right = popList()
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return head

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// JSONArrayToListNode [5,3,6,2,4,null,7] use it
func JSONArrayToListNode(j string) *ListNode {
	arg := JsonStringToSliceAny(j)
	if len(arg) < 1 {
		return nil
	}
	list := make([]*ListNode, 0)

	for _, v := range arg {
		if v == nil {
			list = append(list, nil)
			continue
		}
		tNode := &ListNode{
			Val:  int(v.(float64)),
			Next: nil,
		}
		list = append(list, tNode)
	}

	for idx := range list {
		l := len(list)
		if idx < l-1 {
			list[idx].Next = list[idx+1]
		}
	}

	return list[0]

}
