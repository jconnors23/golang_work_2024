// simulate an elevator software system

/*
features to implement:

fifo queueing, unique vals
timing/goroutines,
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
	floor_selection(&elevator)
	stop()
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
	fmt.Printf("Base Floor: 1\n")
	fmt.Printf("Select floor(s) (0 to stop):\n%d\n ", floors)
	for {
		var button int
		fmt.Scanln(&button)
		if button == 0 {
			break
		} else {
			floor_selections = append(floor_selections, button)
		}
	}
	for _, v := range floor_selections {
		enqueue(e, v)
	}
	ascend_descend(e)
	if len(e.Queue) > 0 {
		floor_selection(e) // recursive if queue has floors still in it, user is also allowed to augment queue with new floor inputs
	}
}

func enqueue(e *Elevator, desired_floor int) {
	to_queue := desired_floor
	if to_queue < 1 || to_queue > 10 {
		stop()
	} else {
		e.Queue = append(e.Queue, to_queue) // enqueue
	}
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
	// each floor will take 10 units of time to ascend / descend to
	trip_time := int(distance * 10)
	for trip_time > 0 {
		time.Sleep(2 * time.Second) // mimic elevator ascending/descending
		fmt.Printf("'Beep'. %v floor", e.Floor)
		trip_time -= 10
	}
	e.Floor = e.Queue[0]  // current floor after we have moved up or down
	e.Queue = e.Queue[1:] // dequeue, remove the floor that we traveled to from our queue
	fmt.Printf("'Ding'. %v floor", e.Floor)
	if len(e.Queue) == 0 {
		stop()
	}
}

func stop() {
	fmt.Println("No floors queued or service unavailable.")
	os.Exit(0)
}
