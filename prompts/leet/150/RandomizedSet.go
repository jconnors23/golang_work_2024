package main

import (
	"math/rand"
	"sort"
)

//https://leetcode.com/problems/insert-delete-getrandom-o1/description/?envType=study-plan-v2&envId=top-interview-150

// stack behavior

func main() {

}

type RandomizedSet struct {
	set []int
}

func Constructor() RandomizedSet {
	new_set := RandomizedSet{
		set: []int{},
	}
	return new_set
}

func (this *RandomizedSet) Insert(val int) bool {
	if val > this.Peep() {
		this.set = append(this.set, val)
		return true
	}
	for _, v := range this.set {
		if v == val {
			return false
		}
	}
	this.set = append(this.set, val)
	return true
}

func (this *RandomizedSet) Peep() int {
	return this.set[0]
}

func (this *RandomizedSet) Stack() {
	sort.Ints(this.set)
}

func (this *RandomizedSet) Remove(val int) bool {
	for index, v := range this.set {
		if v == val {
			this.set = append(this.set[:index], this.set[index+1:]...)
			this.Stack()
			return true
		}
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	randomIndex := rand.Intn(len(this.set))
	return this.set[randomIndex]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
