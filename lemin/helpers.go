package lemin

import (
	"fmt"
	"log"
	"strings"
)

/* To find the rooms after ##start or ##end (useful when comments are in the way)*/
func FindNextNonComment(arr []string, start int) int {
	if start > len(arr) {
		fmt.Println("Error Handling just in case, although im not sure if this will ever activate.")
	}
	for i := start; i < len(arr); i++ {
		if len(arr) < 2 && i != len(arr)-1 {
			log.Fatal("ERROR: Invalid data format. The only possible empty line should be the last line.")
		} else if len(arr) > 2 {
			if len(arr[i]) == 0 {
				log.Fatal("ERROR: Invalid data format. Unexpected empty line found.")
			}
			if arr[i][0] != '#' {
				return i
			}
		}
	}
	return -1
}

/* This function checks if room format is: <room name> <x coordinate> <y coordinate> */
func IsValidRoom(arr []string, index int) {
	temp := strings.Split(arr[index], " ")
	if len(temp) != 3 {
		log.Fatal("ERROR: Invalid room format. Correct usage: <room_name> <coor_x> <coord_y>.")
	}
	if strings.ToLower(temp[0])[0] == 'l' {
		log.Fatal("ERROR: Invalid data format. A room will never start with the letter 'L'.")
	}
	for _, elem := range temp[1] {
		if elem < '0' || elem > '9' {
			log.Fatal("ERROR: Invalid data format. Non-numeric value in coordinates found.")
		}
	}
	for _, elem := range temp[2] {
		if elem < '0' || elem > '9' {
			log.Fatal("ERROR: Invalid data format. Non-numeric value in coordinates found")
		}
	}
}

/* This functions throws an error if two rooms have the same exact coordinates */
func (g *Graph) ValidCoord() {
	for _, elem := range g.Vertices {
		for _, elem2 := range g.Vertices {
			if elem.key != elem2.key && elem.coord_x == elem2.coord_x && elem.coord_y == elem2.coord_y {
				log.Fatalf("ERROR: Invalid data. Two or more vertices have matching coordinates \n[%v and %v]", elem.key, elem2.key)
			}
		}
	}
}
