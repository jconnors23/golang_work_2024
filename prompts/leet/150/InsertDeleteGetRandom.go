package main

import (
	"math/rand"
	"sort"
)

//https://leetcode.com/problems/insert-delete-getrandom-o1/description/?envType=study-plan-v2&envId=top-interview-150

// stack behavior

type RandomizedSet struct {
	set []int
}

func Constructor() RandomizedSet {
	new_set := RandomizedSet{
		set: []int{},
	}
	return new_set
}

func (s *RandomizedSet) Insert(val int) bool {
	for _, v := range s.set {
		if v == val {
			return false
		}
	}
	s.set = append(s.set, val)
	return true
}

func (s *RandomizedSet) Stack() {
	sort.Ints(s.set)
}

func (s *RandomizedSet) Remove(val int) bool {
	for index, v := range s.set {
		if v == val {
			s.set = append(s.set[:index], s.set[index+1:]...)
			s.Stack()
			return true
		}
	}
	return false
}

func (s *RandomizedSet) GetRandom() int {
	randomIndex := rand.Intn(len(s.set))
	return s.set[randomIndex]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
