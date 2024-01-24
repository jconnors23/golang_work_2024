// simulate an elevator software system

/*
features to implement:
able to enter floors always
actual timing
fifo queueing, unique vals
timing/goroutines,
beep funcionality show current floor as we a, d
fmt.Printf("'Beep'. %v floor\n", e.Floor)
firealarm logic, call stop when user pulls firealarm
scaling time for how long elevator would take to reach each floor
print 2d array showing floors in building
queue can have go routine and if empty after 10 seconds os running status becomes false?
logic based on distance bw floors to deque

[9 1 8]
*/
package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"time"
)

func main() {
	elevator := Elevator{Floor: 1}
	fmt.Printf("Base Floor: 1\n")
	floor_selection(&elevator)
	sort_by_distance(&elevator)
	for len(elevator.Queue) > 0 {
		ascend_descend(&elevator)
		// prevent index out of range
		if len(elevator.Queue) == 1 {
			ding(&elevator, elevator.Queue[0])
			floor_selection(&elevator)
			sort_by_distance(&elevator)
			// if no additional floors entered
			if len(elevator.Queue) == 1 {
				elevator.Queue = []int{1}
				ascend_descend(&elevator)
				stop()
			}
		}
		ding(&elevator, elevator.Queue[0])
		elevator.Queue = elevator.Queue[1:] // dequeue, remove floor after travel
	}
}

type Elevator struct {
	Floor int
	Queue []int
}

func floor_selection(e *Elevator) {
	var floor_selections []int
	floors := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("Select floor(s) (0 to stop):\n%d\n ", floors) // needs scanning go routine
	for {
		var button int
		fmt.Scanln(&button)
		if button == 0 {
			break
		} else {
			floor_selections = append(floor_selections, button)
			fmt.Printf("%d\n", floor_selections)
		}
	}
	for _, v := range floor_selections {
		enqueue(e, v)
	}
}

func enqueue(e *Elevator, desired_floor int) {
	to_queue := desired_floor
	if to_queue < 1 || to_queue > 10 {
		stop()
	} else if contains(e, to_queue) {
		return
	} else {
		e.Queue = append(e.Queue, to_queue) // enqueue
	}
}

func contains(e *Elevator, to_queue int) bool {
	for _, floor := range e.Queue {
		if floor == to_queue {
			return true
		}
	}
	return false
}

func ascend_descend(e *Elevator) {
	if e.Queue[0] > e.Floor {
		arrow := "\u2191"
		fmt.Printf("Ascending %v \n", arrow)
	} else if e.Queue[0] < e.Floor {
		arrow := "\u2193"
		fmt.Printf("Descending %v \n", arrow)
	}
	// distance in b/w floors
	distance := absolute_value(e.Queue[0], e.Floor)
	// each floor will take 2 units of time to ascend / descend to
	trip_time := int(distance * 2)
	for trip_time > 0 {
		time.Sleep(time.Second) // mimic elevator ascending/descending
		trip_time -= 2
	}
}

func ding(e *Elevator, current int) {
	e.Floor = current // floor after movement
	fmt.Printf("'Ding'. %v floor\n", current)
}

// sort our queue by distance to current floor
func sort_by_distance(e *Elevator) {
	distace_to := e.Queue[0]
	sort.Slice(e.Queue, func(i, j int) bool {
		return absolute_value(distace_to, e.Queue[i]) < absolute_value(distace_to, e.Queue[j])
	})
	fmt.Println(e.Queue)
}

func absolute_value(x, y int) float64 {
	return math.Abs(float64(x) - float64(y))
}

func stop() {
	fmt.Printf("Base Floor: 1\n")
	fmt.Println("No floors queued or service unavailable.")
	os.Exit(0)
}
