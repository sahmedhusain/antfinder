package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Graph struct {
    Nodes map[string]*Node
    Start *Node
    End   *Node
}

type Node struct {
    Name     string
    Edges    []*Node
}

type Ant struct {
    ID   int
    Path []*Node
}

func ReadFile(filePath string) ([]string, error) {
    file, err := os.Open(filePath)
    if (err != nil) {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return lines, nil
}

func ParseLines(lines []string) ([]string, []string, int, string, string, error) {
    var connections []string
    var nodeDefinitions []string
    var ants int
    var startRoomName, endRoomName string

    for i := 0; i < len(lines); i++ {
        line := lines[i]

        if line == "" {
            continue
        }

        if line == "##start" {
            if i+1 < len(lines) {
                startRoomName = lines[i+1]
                fmt.Println("Found start room:", startRoomName) // Debug print
                i++ // Skip the next line as it is the start room name
            }
            continue
        }

        if line == "##end" {
            if i+1 < len(lines) {
                endRoomName = lines[i+1]
                fmt.Println("Found end room:", endRoomName) // Debug print
                i++ // Skip the next line as it is the end room name
            }
            continue
        }

        if ants == 0 {
            fmt.Sscanf(line, "%d", &ants)
            fmt.Println("Found ants:", ants) // Debug print
        } else if strings.Contains(line, "-") {
            connections = append(connections, line)
        } else {
            nodeDefinitions = append(nodeDefinitions, line)
        }
    }

    fmt.Println("Final start room:", startRoomName) // Debug print
    fmt.Println("Final end room:", endRoomName) // Debug print

    if startRoomName == "" || endRoomName == "" {
        return nil, nil, 0, "", "", fmt.Errorf("invalid data format, no start or end room found")
    }

    return nodeDefinitions, connections, ants, startRoomName, endRoomName, nil
}

func BuildGraph(nodeDefinitions, connections []string, startRoomName, endRoomName string) (*Graph, error) {
    graph := &Graph{
        Nodes: make(map[string]*Node),
    }

    // Create nodes from node definitions
    for _, nodeDef := range nodeDefinitions {
        parts := strings.Fields(nodeDef)
        if len(parts) != 3 {
            return nil, fmt.Errorf("invalid node definition format: %s", nodeDef)
        }
        nodeName := parts[0]
        graph.Nodes[nodeName] = &Node{Name: nodeName}
    }

    // Create edges from connections
    for _, connection := range connections {
        parts := strings.Split(connection, "-")
        if len(parts) != 2 {
            return nil, fmt.Errorf("invalid connection format: %s", connection)
        }
        node1Name, node2Name := parts[0], parts[1]

        node1, exists := graph.Nodes[node1Name]
        if !exists {
            return nil, fmt.Errorf("node not found: %s", node1Name)
        }

        node2, exists := graph.Nodes[node2Name]
        if !exists {
            return nil, fmt.Errorf("node not found: %s", node2Name)
        }

        node1.Edges = append(node1.Edges, node2)
        node2.Edges = append(node2.Edges, node1)
    }

    // Set the start and end nodes
    startNode, exists := graph.Nodes[startRoomName]
    if !exists {
        return nil, fmt.Errorf("start room not found: %s", startRoomName)
    }
    graph.Start = startNode

    endNode, exists := graph.Nodes[endRoomName]
    if !exists {
        return nil, fmt.Errorf("end room not found: %s", endRoomName)
    }
    graph.End = endNode

    return graph, nil
}

func BFS(graph *Graph) []*Node {
    queue := []*Node{graph.Start}
    visited := make(map[string]bool)
    prev := make(map[string]*Node)

    visited[graph.Start.Name] = true

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        if current == graph.End {
            break
        }

        for _, neighbor := range current.Edges {
            if !visited[neighbor.Name] {
                queue = append(queue, neighbor)
                visited[neighbor.Name] = true
                prev[neighbor.Name] = current
            }
        }
    }

    var path []*Node
    for at := graph.End; at != nil; at = prev[at.Name] {
        path = append([]*Node{at}, path...)
    }

    return path
}

func simulateAnts(graph *Graph, path []*Node, ants int) {
    antPositions := make([]int, ants)
    for i := range antPositions {
        antPositions[i] = 0
    }

    for {
        moved := false
        for i := 0; i < ants; i++ {
            if antPositions[i] < len(path)-1 {
                antPositions[i]++
                fmt.Printf("L%d-%s ", i+1, path[antPositions[i]].Name)
                moved = true
            }
        }
        if !moved {
            break
        }
        fmt.Println()
    }
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("ERROR: No input file provided")
        return
    }

    filePath := os.Args[1]
    lines, err := ReadFile(filePath)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    nodeDefinitions, connections, ants, startRoomName, endRoomName, err := ParseLines(lines)
    if err != nil {
        fmt.Println("Error parsing lines:", err)
        return
    }
    graph, err := BuildGraph(nodeDefinitions, connections, startRoomName, endRoomName)
    if err != nil {
        fmt.Println("Error building graph:", err)
        return
    }

    path := BFS(graph)
    simulateAnts(graph, path, ants)
}
