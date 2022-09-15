package q2353designafoodratingsystem

type FoodRatings struct {
	cMaxHeap       map[string]*foodMaxHeap // cuisine map to max heap
	foodCuisineMap map[string]string       // food map to cuisine
}
type food struct {
	name    string
	cuisine string
	rating  int
}
type foodMaxHeap struct {
	fmh        []food
	foodIdxMap map[string]int // food map to index in fmh
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	f := FoodRatings{
		cMaxHeap:       make(map[string]*foodMaxHeap),
		foodCuisineMap: make(map[string]string),
	}
	for idx, name := range foods {
		cuisine := cuisines[idx]
		rating := ratings[idx]
		if _, ok := f.cMaxHeap[cuisine]; !ok {
			x := foodMaxHeap{
				fmh:        make([]food, 0),
				foodIdxMap: make(map[string]int),
			}
			f.cMaxHeap[cuisine] = &x
		}
		f.cMaxHeap[cuisine].insert(food{
			name:    name,
			cuisine: cuisine,
			rating:  rating,
		})

		f.foodCuisineMap[name] = cuisine
	}
	return f
}

func (f *FoodRatings) ChangeRating(foodName string, newRating int) {
	cuisine := f.foodCuisineMap[foodName]
	idx := f.cMaxHeap[cuisine].foodIdxMap[foodName]
	f.cMaxHeap[cuisine].remove(idx)
	f.cMaxHeap[cuisine].insert(food{
		name:    foodName,
		cuisine: cuisine,
		rating:  newRating,
	})
	f.foodCuisineMap[foodName] = cuisine
}

func (f *FoodRatings) HighestRated(cuisine string) string {
	return f.cMaxHeap[cuisine].fmh[0].name
}

// --------------

func (f *foodMaxHeap) insert(fd food) {
	l := len(f.fmh)
	f.fmh = append(f.fmh, fd)
	f.foodIdxMap[fd.name] = l
	f.upperIfy(l)
}
func (f *foodMaxHeap) upperIfy(idx int) {
	pi := f.parent(idx)
	if pi == idx {
		return
	}
	if f.fmh[pi].rating < f.fmh[idx].rating {
		f.swap(pi, idx)
		f.upperIfy(pi)
	} else if f.fmh[pi].rating == f.fmh[idx].rating {
		if f.fmh[pi].name > f.fmh[idx].name {
			f.swap(pi, idx)
		}
	}
}
func (f *foodMaxHeap) downerIfy(idx int) {
	maxIdx := len(f.fmh) - 1
	l := f.left(idx)
	r := f.right(idx)
	compareIdx := 0
	if l > maxIdx {
		return // 最底拉
	} else if r > maxIdx {
		compareIdx = l
	} else if f.fmh[l].rating < f.fmh[r].rating {
		compareIdx = r
	} else {
		compareIdx = l
	}

	if f.fmh[idx].rating < f.fmh[compareIdx].rating {
		f.swap(idx, compareIdx)
		f.downerIfy(compareIdx)
	} else if f.fmh[idx].rating == f.fmh[compareIdx].rating {
		if f.fmh[idx].name > f.fmh[compareIdx].name {
			f.swap(idx, compareIdx)
		}
	}
}
func (f *foodMaxHeap) swap(i1 int, i2 int) {
	food1 := f.fmh[i1]
	food2 := f.fmh[i2]

	f.foodIdxMap[food1.name], f.foodIdxMap[food2.name] = i2, i1
	f.fmh[i1], f.fmh[i2] = f.fmh[i2], f.fmh[i1]
}
func (f *foodMaxHeap) parent(idx int) int {
	n := (idx - 1) / 2
	return n
}
func (f *foodMaxHeap) left(idx int) int {
	return idx*2 + 1
}
func (f *foodMaxHeap) right(idx int) int {
	return idx*2 + 2
}

func (f *foodMaxHeap) remove(idx int) {
	l := len(f.fmh) - 1
	if l != idx {
		// remove foodIdMap
		delete(f.foodIdxMap, f.fmh[idx].name)
		delete(f.foodIdxMap, f.fmh[l].name)

		// put last to current
		f.fmh[idx] = f.fmh[l]
		f.foodIdxMap[f.fmh[idx].name] = idx

		// let current upperIfy
		f.fmh = f.fmh[:l]
		f.upperIfy(idx)
		f.downerIfy(idx)
		return
	}
	delete(f.foodIdxMap, f.fmh[idx].name)
	f.fmh = f.fmh[:l]
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
