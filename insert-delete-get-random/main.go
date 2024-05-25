package main

import (
	"fmt"
	"math/rand/v2"
)

type RandomizedSet struct {
    data map[int]int 
	idx []int
}


func Constructor() RandomizedSet {
    s := RandomizedSet {
		data: make(map[int]int),
		idx: make([]int, 0),
	}

	return s
}


func (this *RandomizedSet) Insert(val int) bool {
    _, ok := this.data[val]
	
	if !ok {
		this.data[val] = len(this.idx)
		this.idx = append(this.idx, val)
		return true
	}else{
		return false
	}

}


func (this *RandomizedSet) Remove(val int) bool {
    idx, ok := this.data[val]
	
	if ok {
		size := len(this.idx)
		this.data[this.idx[size-1]] = idx
		this.idx[size-1], this.idx[idx] = this.idx[idx], this.idx[size-1]
		this.idx = this.idx[:size-1] 
		delete(this.data, val)
		return true
	}else{
		return false
	}
}


func (this *RandomizedSet) GetRandom() int {
	size := len(this.idx)
    rand_idx := rand.IntN(size)

	return this.idx[rand_idx]
}


func main(){
	obj := Constructor()

	obj.Remove(0)
	obj.Remove(0)
	obj.Insert(0)
	fmt.Println(obj.GetRandom())
	obj.Remove(0)
	res := obj.Insert(0)
	fmt.Println(res)
	fmt.Println(obj)


}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */