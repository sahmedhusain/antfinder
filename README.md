# AntFinder 🐜

[![Go](https://img.shields.io/badge/Go-1.23.2-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE.md)
[![Algorithms](https://img.shields.io/badge/Algorithms-Graph%20Traversal-blueviolet)](#-algorithm-explanation)
[![DFS](https://img.shields.io/badge/Algorithm-DFS-orange)](#-pathfinding-logic)
[![Terminal](https://img.shields.io/badge/Interface-Terminal-black)](#-terminal-examples)

Welcome to **AntFinder**, a terminal-based application built with **Go** that simulates efficient pathfinding for ant colonies. The program discovers optimal routes for multiple ants moving through a graph and displays their movements step by step in real time.

AntFinder is designed to be both educational and practical, focusing on algorithmic problem-solving, graph traversal, and movement simulation.

---

## ✨ Features

- **Optimized Pathfinding** 🛤️  
  Finds the shortest paths for ants using graph traversal algorithms.

- **Concurrent Ant Movement** 🐜  
  Simulates multiple ants moving simultaneously without conflicts.

- **Large Input Handling** 📊  
  Processes large graphs and high ant counts efficiently.

- **Error Validation** ⚠️  
  Provides clear error messages for invalid inputs.

- **Terminal-Based Output** 🖥️  
  Displays ant movements in real time via the command line.

---

## 🛠️ Technologies Used

- **Go** 🐹 – Backend processing, algorithm implementation, and data handling.

## 🔧 Technologies & Algorithms

<p align="left">
  <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" width="40" alt="Go"/>
  <img src="https://img.shields.io/badge/DFS-Depth%20First%20Search-orange" alt="DFS"/>
  <img src="https://img.shields.io/badge/Graph-Algorithms-blueviolet" alt="Graph Algorithms"/>
  <img src="https://img.shields.io/badge/Data%20Structures-Graphs%20%26%20Paths-green" alt="Data Structures"/>
  <img src="https://img.shields.io/badge/Interface-Terminal-black" alt="Terminal"/>
</p>

This project combines **Go**, classic **graph algorithms**, and **Depth-First Search (DFS)** to solve an optimization and simulation problem in a terminal-based environment.

---

## 🎯 What We Aim For

AntFinder processes input files describing graphs composed of rooms and tunnels in order to simulate ant movement.

The input includes:

1. **Ants** 🐜 – Number of ants to move  
2. **Rooms** 🏠 – Nodes in the graph, including start and end  
3. **Tunnels** 🚇 – Connections between rooms  

Efficient algorithms are used to find paths and simulate movement so that ants reach the end room with the minimum number of steps.

---

## 📄 Input Format Details

The input file must follow a specific format:

- **First line**: Number of ants (e.g. `5`)
- **Room lines**: `name x y` (coordinates are ignored for pathfinding)
- **Special rooms**:
  - `##start` followed by the start room
  - `##end` followed by the end room
- **Tunnel lines**: Connections between rooms (e.g. `room1-room2`)

### Example Input File

```
5
##start
start 0 0
room1 1 1
room2 2 2
##end
end 3 3
start-room1
room1-room2
room2-end
```

This represents 5 ants moving from `start` to `end` via `room1` and `room2`.

### Input Summary Table

| Line Type    | Example       | Description                               |
|--------------|---------------|-------------------------------------------|
| Ant Count    | `5`           | Number of ants                            |
| Start Marker | `##start`     | Indicates the next line is the start room |
| Room         | `start 0 0`   | Room name and coordinates                 |
| End Marker   | `##end`       | Indicates the next line is the end room   |
| Tunnel       | `start-room1` | Connection between two rooms              |

---

## 🧭 Graph Representation

```
start ---- room1 ---- room2 ---- end
```

- **Nodes** represent rooms  
- **Edges** represent tunnels  

Ants must move from start to end without collisions, using multiple paths when available.

---

## 🚀 Getting Started

### Prerequisites

- Go **1.23.2** or higher installed on your machine.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/antfinder.git
   ```
2. Navigate to the project directory:
   ```bash
   cd antfinder
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run the application:
   ```bash
   go run main.go examples/example00.txt
   ```

---

## 📖 How to Use

Once running, the program reads the input file and outputs ant movements turn by turn. Each line shows the position of ants at that step. The simulation ends when all ants reach the end room.

### Step-by-Step Usage

1. Prepare an input file in the required format  
2. Run the program using `go run main.go path/to/input.txt`  
3. Observe the printed ant movements  
4. Each turn shows how ants advance toward the end room  

---

## 🧠 Algorithm Explanation

---

AntFinder uses graph algorithms to find optimal paths for ant movement.

### Main Steps

1. **Parse Input** – Read and validate the file  
2. **Find All Paths** – Use DFS to discover paths from start to end  
3. **Select Best Paths** – Choose non-overlapping paths  
4. **Simulate Movement** – Move ants turn by turn, avoiding conflicts  

---

### Algorithm Flowchart

![Algorithm Flowchart](flowchart.png)

_The flowchart illustrates the main steps of the AntFinder algorithm, from parsing input to simulating movement, with error-handling branches._

---

### Data Structure ERD

![Data Structure ERD](erd.png)

_The ERD shows relationships between Room, Tunnel, Ant, and Path entities, including their attributes._

---

### Pathfinding Logic

🧠 **Algorithm Used:** Depth-First Search (DFS)

Depth-First Search (DFS) is used to explore all possible paths.

For a simple graph:

```
start -- a -- b -- end
  |         |
  -- c -- d --
```

DFS explores:

- start → a → b → end  
- start → a → d → end  
- start → c → d → end  

All valid paths are collected and evaluated.

---

### Ant Assignment

Ants are assigned to paths to balance load:

| Path       | Length | Ants Assigned |
|------------|--------|---------------|
| P1 (short) | 3      | 3 ants        |
| P2 (long)  | 5      | 2 ants        |

This minimizes the total number of moves.

---

### Movement Simulation

Ants move simultaneously. In each turn:

- Check if the next room is free  
- Move ants if possible  

Example output:

```
L1-a L2-a
L1-b L2-b L3-a
L1-end L2-end L3-b
L3-end
```

---

## 💻 Terminal Examples

---

### Running with Example Input

```bash
$ go run main.go examples/example00.txt
L1-0 L1-1
L1-2 L1-3
```

### Ant Movement Simulation

```bash
$ go run main.go examples/example01.txt
L1-0 L2-0
L1-1 L2-1
L1-2 L2-2
L1-3 L2-3
```

---

## 🛠️ Under the Hood

---

### Data Handling

- **Room Struct** – Name, coordinates, connections  
- **Ant Struct** – ID, current path, position  
- **Graph Struct** – All rooms and tunnels  

### Code Structure

- `main.go` – Entry point  
- `functions/ant.go` – Ant movement logic  
- `functions/utils.go` – Utility functions and validation  
- `datastruct/structs.go` – Core data structures  

### Error Management

Handled cases include:
- Invalid ant count  
- Duplicate rooms or tunnels  
- Missing start or end rooms  
- Disconnected graphs  

Each error prints a clear message before exiting.

---

## 🤝 Contributing

Contributions are welcome. Fork the repository, make changes, and submit a pull request following Go best practices.

---

## 📄 License

Licensed under the **MIT License**. See [LICENSE.md](LICENSE.md) for details.

---

## 🙏 Acknowledgments

This project was created during a Go learning journey, emphasizing algorithm implementation and simulation. Inspired by the classic **lem-in** problem.

---

## 👥 Authors

- **Sayed Ahmed Husain** – [sayedahmed97.sad@gmail.com](mailto:sayedahmed97.sad@gmail.com)  
- **Qasim Aljaffer**  
- **Mohammed AlAlawi**  
- **Abdulla Alasmawi**

---

## 📚 What I Learned

- Graph algorithms and pathfinding techniques  
- Efficient data structures in Go  
- Parsing and validating input files  
- Simulation of concurrent processes  

---

## ⚠️ Limitations

- Input must follow a strict format  
- Maximum supported ants: **10,000**

---

## 🔮 Future Improvements

- Support for dynamic graphs  
- Performance enhancements  
- Optional GUI visualization  
