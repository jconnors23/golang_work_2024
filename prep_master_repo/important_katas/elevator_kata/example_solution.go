/*
 Open ended design for this Kata. This design features:
- queue for floors
- elevator goes to closest relative floor
- direction arrows, UI features, and user prompts
- os exit
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
	floor_selection(&elevator)  // user inputs floors
	sort_by_distance(&elevator) // sort floors based on distance from elevator current floor
	for len(elevator.Queue) > 0 {
		ascend_descend(&elevator)
		// prevent index out of range
		if len(elevator.Queue) == 1 {
			ding(&elevator, elevator.Queue[0])
			floor_selection(&elevator) // user can input more floors
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

// elevator will have current floor and Queue of floors attributes
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
	fmt.Printf("Select floor(s) (0 to stop):\n%d\n ", floors)
	// while loop, when the user selects 0 loop breaks, otherwise floors are added to slice
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
	// enter each floor into the queue using helper method
	for _, v := range floor_selections {
		enqueue(e, v)
	}
}

func enqueue(e *Elevator, desired_floor int) {
	to_queue := desired_floor // var is here to enhance readability
	if to_queue < 1 || to_queue > 10 {
		stop()
	} else if contains(e, to_queue) {
		return
	} else {
		e.Queue = append(e.Queue, to_queue) // enqueue
	}
}

// helper function to simulate if elevator floor has been chosen already
func contains(e *Elevator, to_queue int) bool {
	for _, floor := range e.Queue {
		if floor == to_queue {
			return true
		}
	}
	return false
}

func ascend_descend(e *Elevator) {
	// print ascend, descend arrows
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
		time.Sleep(time.Second) // mimic elevator travel time using time.Sleep
		trip_time -= 2
	}
}

// ui
func ding(e *Elevator, current int) {
	e.Floor = current // floor after movement
	fmt.Printf("'Ding'. %v floor\n", current)
}

// sort our queue based on distance to current floor
func sort_by_distance(e *Elevator) {
	distance_to := e.Queue[0]
	sort.Slice(e.Queue, func(i, j int) bool {
		return absolute_value(distance_to, e.Queue[i]) < absolute_value(distance_to, e.Queue[j])
	})
	fmt.Println(e.Queue)
}

func absolute_value(x, y int) float64 {
	return math.Abs(float64(x) - float64(y))
}

// stop the elevator and exit the program
func stop() {
	fmt.Printf("Base Floor: 1\n")
	fmt.Println("No floors queued or service unavailable.")
	os.Exit(0)
}
