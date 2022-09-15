package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"testing"

	q2349designanumbercontainersystem "github.com/limiu82214/GoLeetcode/exam/Q2349DesignANumberContainerSystem"
	"github.com/stretchr/testify/assert"
)

// TestForCopy ============= 複製一個copy並開始你的表演
func TestForCopy(t *testing.T) {

	type param struct {
		Arg1     []int
		Excepted []int
	}
	data := []param{
		{[]int{1, 1, 2, 2, 2, 3}, []int{3, 1, 1, 2, 2, 2}},
		{[]int{2, 3, 1, 3, 2}, []int{1, 3, 3, 2, 2}},
		{[]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}, []int{5, -1, 4, 4, -6, -6, 1, 1, 1}},
	}

	decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Excepted, Q1636SortArrayByIncreasingFrequency(d.Arg1), d)
		}
	}, 1, 0)

}

// TestQ2353DesignAFoodRatingSystem Medium
func TestQ2353DesignAFoodRatingSystem(t *testing.T) {

}
func TestQ1636SortArrayByIncreasingFrequency(t *testing.T) {
	type param struct {
		Arg1     []int
		Excepted []int
	}
	data := []param{
		{[]int{1, 1, 2, 2, 2, 3}, []int{3, 1, 1, 2, 2, 2}},
		{[]int{2, 3, 1, 3, 2}, []int{1, 3, 3, 2, 2}},
		{[]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}, []int{5, -1, 4, 4, -6, -6, 1, 1, 1}},
	}

	decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Excepted, Q1636SortArrayByIncreasingFrequency(d.Arg1), d)
		}
	}, 1, 0)
}
func TestQ595BigCountries(t *testing.T) {
	// skip
}
func TestQ237DeleteNodeInALinkedList(t *testing.T) {
	// skip
}

func TestQ2349DesignANumberContainerSystem(t *testing.T) {
	type param struct {
		Arg1     string
		Arg2     []int
		Excepted interface{}
	}
	data := []param{
		// {"NumberContainers", []int{}, nil},
		// 1
		// {"find", []int{10}, -1},
		// {"change", []int{2, 10}, nil},
		// {"change", []int{1, 10}, nil},
		// {"change", []int{3, 10}, nil},
		// {"change", []int{5, 10}, nil},
		// {"find", []int{10}, 1},
		// {"change", []int{1, 20}, nil},
		// {"find", []int{10}, 2},

		// 2
		{"change", []int{1, 10}, nil},
		{"find", []int{10}, 1},
		{"change", []int{1, 20}, nil},
		{"find", []int{10}, -1},
		{"find", []int{20}, 1},
		{"find", []int{30}, -1},
	}

	decorate(func() {
		obj := q2349designanumbercontainersystem.Constructor()
		for _, d := range data {
			fmt.Println(d, obj)
			switch d.Arg1 {
			case "find":
				assert.Equal(t, d.Excepted, obj.Find(d.Arg2[0]))
			case "change":
				obj.Change(d.Arg2[0], d.Arg2[1])
				assert.Equal(t, d.Excepted, nil)
			}
		}
	}, 1, 0)
}

func TestQ1TwoSum(t *testing.T) {
	type param struct {
		Arg1     []int
		Arg2     int
		Excepted []int
	}
	data := []param{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	decorate(func() {
		for _, d := range data {
			nums := d.Arg1
			target := d.Arg2
			actual := Q1TwoSum(nums, target)
			excepted := d.Excepted
			assert.Equal(t, excepted, actual, nums)
		}
	}, 3000, 10000)
}

func TestQ2180CountIntegersWithEvenDigitSum(t *testing.T) {
	type param struct {
		Arg1     int
		Excepted int
	}
	data := []param{
		{4, 2},
		{30, 14},
	}

	decorate(func() {
		for _, d := range data {
			actual := Q2180CountIntegersWithEvenDigitSum(d.Arg1)
			excepted := d.Excepted
			assert.Equal(t, excepted, actual, d)
		}
	}, 1, 0)
}

func decorate(inter func(), times int, cpurate int) {
	if cpurate == 0 {
		cpurate = 100
	}
	runtime.SetCPUProfileRate(cpurate)
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
