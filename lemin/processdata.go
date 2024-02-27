package lemin

// This file contains all data processing functions that occur in the bigenning of the program.
import (
	"log"
	"strconv"
	"strings"
)

var Farm = Graph{} // Global variable for easy use

/* This function processes all data in data file, and initializes all the data into the farm graph */
func ProcessData(d []string) {
	var (
		foundStart  = false
		foundEnd    = false
		foundAnts   = false
		indexOfAnts = 0
	)

	/* Loop for finding number of ants */
	for i, line := range d {
		ant, err := strconv.Atoi(line) // might need a variable name change here

		// if there is no error, that means there is no spaces or alphabetical characters -> It is a number.
		if err == nil && foundAnts {
			log.Fatal("ERROR: Invalid data format. Number of ants should be mentioned only once before the start.")
		} else if err == nil {
			Farm.ants = ant
			foundAnts = true
			indexOfAnts = i
			continue
		}
		// If start or end is found before num of ants, then the data format is invalid.
		temp := strings.ToLower(line)
		if temp == "##start" || temp == "##end" {
			if !foundAnts {
				log.Fatal("ERROR: Invalid data format. Start/end room found before number of ants.")
			} else {
				break // If start is found and ants is also found, then break out of the loop.
			}
		}
	}

	// Loop for graph structuring
	for i, line := range d {
		// if comment (single #), continue...
		if len(line) > 2 && line[0] == '#' && line[1] != '#' {
			continue
		}

		// find start and finish
		if len(line) > 2 && line[0] == '#' && line[1] == '#' {
			if len(d)-2 < i { // Checking if start or end command is at the end of file.
				log.Fatal("ERROR: Invalid data format. Start or end room is followed by nothing.")
			}
			if strings.ToLower(line[2:]) == "start" { // converting to lower for case insensitivity
				if foundStart {
					log.Fatal("ERROR: Invalid data format. Start command is mentioned twice.")
				} else {
					foundStart = true
				}

				// check if ##end immediately follows start, if so -> return err
				if d[i+1] == "##end" || d[i-1] == "##end" {
					log.Fatal("ERROR: Start and end are the same room") // error start and end are the same room
				} else {
					// add start room
					start := FindNextNonComment(d, i)
					if start == -1 {
						log.Fatal("ERROR: Invalid data format. start command found but start room not found.")
					}
					IsValidRoom(d, start) //Checking room format
					Farm.startRoom = strings.Split(d[start], " ")[0]
					continue
				}

			} else if strings.ToLower(line[2:]) == "end" { // converting to lower for case insensitivity
				if foundEnd {
					log.Fatal("ERROR: Invalid data format. End command is mentioned twice.")
				} else {
					foundEnd = true
				}
				// add end room
				end := FindNextNonComment(d, i)
				if end == -1 {
					log.Fatal("ERROR: Invalid data format. End command found but end room not found.")
				}
				IsValidRoom(d, end) // Checking room format
				Farm.endRoom = strings.Split(d[end], " ")[0]
				continue
			} else {
				continue // According to the question requirements, all unknown commands should be ignored (in example: ##STOP, ##Hello etc.)
			}
		}

		// split lines to determine which lines are room and which ones are connections
		// if len(splitline) > 1, line has room name, x coord, y coord
		// else if len(splitline) = 1 and not empty, line is connection
		splitLine := strings.Split(line, " ")
		if len(splitLine) > 1 { // room
			IsValidRoom(d, i)
			err := Farm.Addvertex(splitLine[0], splitLine[1], splitLine[2])
			if err != nil {
				log.Fatal(err)
			}
		} else if i != indexOfAnts && splitLine[0] != "" { // connection
			//split on '-' to get from and to connections
			splitCon := strings.Split(strings.Join(splitLine, ""), "-")
			if len(splitCon) != 2 {
				log.Fatal("ERROR: Invalid data format. Connection should be between two rooms and two rooms only.")
			}
			from, to := splitCon[0], splitCon[1]
			if from == to {
				log.Fatal("ERROR: Invalid data. A room must not be connected to itself.")
			}
			//create vertex connection
			Farm.AddConnection(from, to)
		}
	}
	// Farm.Print()
}
