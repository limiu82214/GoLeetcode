package main

import (
	"encoding/json"
	"log"
	"os"
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
	for i := 0; i < times; i++ {
		inter()
	}
}

func JsonStringToSlice(s string) []interface{} {
	var slice []interface{}
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

// JSONArrayToTreeNode [5,3,6,2,4,null,7] use it
func JSONArrayToTreeNode(j string) *TreeNode {
	arg := JsonStringToSlice(j)
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
