package lemin

import (
	"errors"
	"fmt"
	"log"
	"slices"
)

type Graph struct {
	Vertices  []*vertex
	startRoom string
	endRoom   string
	ants      int
}
type vertex struct {
	key         string
	coord_x     string
	coord_y     string
	Connections []*vertex
}

func (g *Graph) Addvertex(key string, x, y string) error {
	if g.Contains(key) {
		err := fmt.Sprintf("vertex %s already exists", key)
		return errors.New(err)
	}
	g.Vertices = append(g.Vertices, &vertex{key: key, coord_x: x, coord_y: y})
	return nil
}

func (g *Graph) AddConnection(from, to string) {
	Vfrom, err := g.GetVertex(from)
	if err != nil {
		log.Fatal(err)
	}
	Vto, err := g.GetVertex(to)
	if err != nil {
		log.Fatal(err)
	}
	if slices.Contains(Vfrom.Connections, Vto) || slices.Contains(Vto.Connections, Vfrom) {
		log.Fatal("ERROR: Two rooms can't be connected with more than two lines.")
	}
	Vfrom.Connections = append(Vfrom.Connections, Vto)
	Vto.Connections = append(Vto.Connections, Vfrom)
}

func (g *Graph) GetVertex(key string) (*vertex, error) {
	for _, v := range g.Vertices {
		if v.key == key {
			return v, nil
		}
	}
	return nil, errors.New("no vertex exists")
}

func (g *Graph) Contains(key string) bool {
	for _, v := range g.Vertices {
		if v.key == key {
			return true
		}
	}
	return false
}

/* This method returns a replicate graph, with the parameter node removed from the replicate's vertices and connections. */
func (g *Graph) Remove(node *vertex) *Graph {
	if node.key == g.startRoom {
		log.Fatal("ERROR! Start room should not be deleted.")
	}

	thisGraph := g.Replicate()

	for i := 0; i < len(thisGraph.Vertices); {
		if thisGraph.Vertices[i].key == node.key {
			thisGraph.Vertices = append(thisGraph.Vertices[:i], thisGraph.Vertices[i+1:]...)
			continue
		}
		var newCon []*vertex
		for _, con := range thisGraph.Vertices[i].Connections {
			if con.key != node.key {
				newCon = append(newCon, con)
			}
		}
		thisGraph.Vertices[i].Connections = newCon
		i++
	}

	return thisGraph
}

/* 
	This function creates a replicate graph and returns it. So far it is only used in the delete method.
	The purpose of replicating is to be able to modify a graph without accidentally modifying the parent graph beacuse of pointer memory addresses.
*/
func (g *Graph) Replicate() *Graph {
	newGraph := &Graph{
		Vertices:  make([]*vertex, len(g.Vertices)),
		startRoom: g.startRoom,
		endRoom:   g.endRoom,
		ants:      g.ants,
	}

	// Copy vertices
	for i, v := range g.Vertices {
		newVertex := &vertex{
			key:         v.key,
			coord_x:     v.coord_x,
			coord_y:     v.coord_y,
			Connections: make([]*vertex, len(v.Connections)),
		}
		newGraph.Vertices[i] = newVertex
	}

	// Copy connections
	for i, v := range g.Vertices {
		for j, con := range v.Connections {
			connection, err := newGraph.GetVertex(con.key)
			if err != nil {
				log.Fatal(err)
			}
			newGraph.Vertices[i].Connections[j] = connection
		}
	}
	return newGraph
}

func (g *Graph) Print() {
	fmt.Println("Ants:", g.ants, "Start:", g.startRoom, "End:", g.endRoom)
	for _, v := range g.Vertices {
		fmt.Printf("vertex key:" + v.key)
		fmt.Print("\tIts connections: ")
		for _, connection := range v.Connections {
			fmt.Print(connection.key, " ")
		}
		fmt.Println()
	}
}
