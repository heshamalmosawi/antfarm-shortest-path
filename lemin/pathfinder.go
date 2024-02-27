package lemin

import (
	"fmt"
	"log"
)

func (g *Graph) PathFinder() {
	paths := make([][]string, 0)
	startVertex := g.findStartVertex()
	if startVertex == nil {
		log.Fatal("Start room not found")
	}
	g.findPaths(startVertex, []string{g.startRoom}, &paths)
	// for _, path := range paths {
	// 	fmt.Printf("Path through room: %v ==> Path:%v\n", path[1], path) // Temporary check for paths
	// }
	paths = QuickSort(paths)
	fmt.Println("\n------------ After sorting ------------")
	for _, path := range paths {
		fmt.Printf("Path through room: %v ==> Path:%v\n", path[1], path) // Temporary check for paths
	}
	fmt.Println("\n------------ After Disjoint ------------")
	paths = disjointPaths(paths)
	for _, path := range paths {
		fmt.Printf("Path through room: %v ==> Path:%v\n", path[1], path) // Temporary check for paths
	}
	fmt.Println()
}

// findStartVertex returns the vertex for the start room.
func (g *Graph) findStartVertex() *vertex {
	for _, v := range g.Vertices {
		if v.key == g.startRoom {
			return v
		}
	}
	return nil
}

// findPaths recursively finds all paths from the current vertex.
func (g *Graph) findPaths(v *vertex, path []string, paths *[][]string) {
	if v.key == g.endRoom {
		fmt.Printf("Path through room: %v ==> Path:%v\n", path[1], path) // Temporary check for paths
		*paths = append(*paths, path)
		return
	}
	for _, conn := range v.Connections {
		// Avoid revisiting vertices in the same path to prevent cycles.
		if !contains(path, conn.key) {
			newPath := append([]string(nil), path...) // Make a copy of the path
			newPath = append(newPath, conn.key)
			g.findPaths(conn, newPath, paths)
		}
	}
}

// contains checks if a slice contains a string.
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func FindShortestPath(paths [][]string) [][]string {
	var sPaths [][]string
	for _, p := range paths {
		var tmp [][]string
		room := p[1]
		tmp = append(tmp, p)
		for _, p2 := range paths {
			if p2[1] == room {
				tmp = append(tmp, p2)
			}
		}
		// fmt.Printf("\nat index %d path:\n %v\n\n", i, tmp)
		sPaths = append(sPaths, findShortest(tmp))
	}
	return sPaths
}
func findShortest(paths [][]string) []string {
	var s []string
	s = paths[0]
	for _, p := range paths {
		if len(p) < len(s) {
			s = p
		}
	}
	return s
}
