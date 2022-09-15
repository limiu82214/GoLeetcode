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
