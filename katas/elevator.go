// simulate an elevator software system

/*
features to implement:

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
	"time"
)

func main() {
	elevator := Elevator{Floor: 1}
	fmt.Printf("Base Floor: 1\n")
	for len(elevator.Queue) >= 1 {
		floor_selection(&elevator)
		ascend_descend(&elevator)
		if len(elevator.Queue) == 1 {

		}
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
	// if len(e.Queue) > 0 {
	// 	floor_selection(e) // recursive if queue has floors still in it, user is also allowed to augment queue with new floor inputs
	// }
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
	distance := math.Abs(float64(e.Queue[0]) - float64(e.Floor))
	// each floor will take 2 units of time to ascend / descend to
	trip_time := int(distance * 2)
	for trip_time > 0 {
		time.Sleep(time.Second) // mimic elevator ascending/descending
		trip_time -= 2
	}
	// prevent index out of range
	if len(e.Queue) == 1 {
		e.Floor = e.Queue[0] // current floor after we have moved up or down
		fmt.Printf("'Ding'. %v floor\n", e.Floor)
		floor_selection(e)
		// if no additional floors entered
		if len(e.Queue) == 1 {
			stop()
		} else {
			return
		}

	}
	e.Queue = e.Queue[1:] // dequeue, remove the floor that we traveled to from our queue
	e.Floor = e.Queue[0]  // current floor after we have moved up or down
	fmt.Printf("'Ding'. %v floor\n", e.Floor)
}

func stop() {
	fmt.Println("No floors queued or service unavailable.")
	os.Exit(0)
}
