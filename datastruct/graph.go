package structs

import (
	"fmt"
	"strings"
)

func NewGraph() *Graph {
	return &Graph{
		Rooms: make(map[string]*Room),
	}
}

func (g *Graph) AddRoom(name string) error {
	if _, exists := g.Rooms[name]; exists {
		return fmt.Errorf("room %s already exists", name)
	}
	g.Rooms[name] = &Room{
		Name:        name,
		Connections: make(map[string]*Room),
	}
	return nil
}

func (g *Graph) AddConnection(room1, room2 string) error {
	r1, r1Exists := g.Rooms[room1]
	r2, r2Exists := g.Rooms[room2]
	if !r1Exists || !r2Exists {
		return fmt.Errorf("one or both rooms do not exist")
	}
	r1.Connections[room2] = r2
	r2.Connections[room1] = r1
	return nil
}

func BuildGraph(connections map[string][]string) *Graph {
	graph := NewGraph()
	for room, links := range connections {
		graph.AddRoom(room)
		for _, link := range links {
			graph.AddRoom(link)
			graph.AddConnection(room, link)
		}
	}
	return graph
}

func ParseLines(lines []string) map[string][]string {
	connections := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			room1 := parts[0]
			room2 := parts[1]
			connections[room1] = append(connections[room1], room2)
			connections[room2] = append(connections[room2], room1)
		}
	}
	return connections
}
