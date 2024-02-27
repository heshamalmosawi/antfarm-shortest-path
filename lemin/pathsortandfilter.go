package lemin

/*
	This function performs the Quick Sort Algorithm on the two-dimensional array consisting of paths,

and sorts them in ascending order based on path length.
*/
func QuickSort(paths [][]string) [][]string {
	// stopping condition for the end of the recursion, if the length is 0 or 1, it will return itself.
	if len(paths) <= 1 {
		return paths
	}

	// Farm.Print()
	// calling the partition function that would put the pivot in its place and return its index for partitioning/dividing the recursion input
	pivot := partition(paths)
	// partitioning/dividing the recursion input arrays
	lower := QuickSort(paths[:pivot])
	higher := QuickSort(paths[pivot+1:])
	// var answer [][]string
	// answer = append(answer, lower...)
	// answer = append(answer, paths[pivot])
	// answer = append(answer, higher...)
	// return answer
	// Farm.test()
	return append(append(lower, paths[pivot]), higher...) // combining all in one line for shorter code
}

/*
	This function selects a pivot element, places the elements with lower length than the pivot on the left side of the pivot and

the elements with bigger length on the right side.
*/
func partition(paths [][]string) int {
	pivot := len(paths[0])  // choosing first element as the 'pivot'
	i, j := 1, len(paths)-1 // saving the index of the first and last element to use as "pointers" (not really pointers)

	for i <= j { // will continue until the first and second point (i, j) cross paths in the array, where it will be i>j
		for i <= j && len(paths[i]) <= pivot {
			// Iterate over all paths that their length are smaller or equal to the pivot, then it will stop when pivot > length of path
			i++
		}
		for i <= j && len(paths[j]) >= pivot {
			// Iterate over all paths with length larger or equal to the pivot, then stop when pivot < path length.
			j--
		}
		if i <= j { // if i & j cross paths, swap their contents
			paths[i], paths[j] = paths[j], paths[i]
		}
	}
	// swap the pivot to its place (currently index j)
	paths[0], paths[j] = paths[j], paths[0]
	return j // return index j
}

/* This function finds the disjoint paths, with priority to the shortest path(s).*/
func disjointPaths(paths [][]string) [][]string {
	roomsUsed := make(map[string]bool)  // Map to store the rooms used so far, for effeciency
	disjointPath := make([][]string, 0) // Collecting the disjoint paths

	for _, path := range paths {
		validPath := true
		for _, room := range path {
			if room == Farm.startRoom || room == Farm.endRoom { // ignore start and end rooms
				continue
			}
			if roomsUsed[room] { // if does not exist in roomsUsed, then it is false
				validPath = false // invalid because room already used
				break
			}
			roomsUsed[room] = true // registering the room into the map
		}
		if validPath {
			disjointPath = append(disjointPath, path)
		}
	}

	// if the length of paths is 0, then just return the first path, this is just in case of anything going wrong, hopefully will never activate.
	if len(disjointPath) == 0 {
		return [][]string{paths[0]}
	}
	return disjointPath
}

func (g *Graph) test(){
	// x, _ := g.GetVertex(g.startRoom)
	// for _, node := range x.Connections{
	// 	graph2 := g.Remove(node)
	// 	fmt.Println(node.key)
	// 	graph2.Print()
	// // os.Exit(1)
	// }
	
}