# lem-in

## Features
- Efficient pathfinding for multiple ants in a graph.
- Handles large input sizes with optimized algorithms.
- Provides detailed error messages for invalid input formats.

## Getting Started

### Prerequisites
- Go 1.23.2 or higher

### Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/lem-in.git
    ```
2. Navigate to the project directory:
    ```sh
    cd lem-in
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```

## Example Usage
To run the project with an example file:
```sh
go run main.go example00.txt
```

## Algorithms and Code Explanation

### Pathfinding and Ant Movement
The core of the project involves finding the quickest path for ants to move from the start room to the end room. The algorithms are implemented in various functions within the `functions/utils.go` and `functions/ant.go` files.

#### Key Functions:
- `ProcessFile`: Reads the input file and initializes the state.
- `startPathOpt`: Recursively finds all possible paths from the start to the end room.
- `Antsop`: Manages the movement of ants along the discovered paths, ensuring the shortest path is taken and avoiding traffic jams.

#### Utility Functions:
- `isUniqueStringSet`: Checks if a set of strings contains unique elements.
- `reverseHyphenatedString`: Reverses the order of hyphen-separated strings.
- `contains`: Checks if a slice contains a specific item.
- `sortByLength`: Sorts slices of strings by their length.
- `equalizeSlices`: Ensures all slices have the same length by padding with "wait".

### Error Handling
The program includes robust error handling to manage various invalid input scenarios, such as:
- Invalid number of ants.
- Missing start or end rooms.
- Duplicate rooms or tunnels.
- Invalid room coordinates.

Errors are handled gracefully with specific error messages, ensuring the program does not crash unexpectedly.

## Limitations
- The program assumes the input format is strictly followed.
- The maximum number of ants is limited to 10,000.

## Future Improvements
- Add support for more complex graph structures.
- Improve the performance of the pathfinding algorithm.
- Add more comprehensive error handling and validation.

## Contributing
1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/NewFeature`).
3. Commit your changes (`git commit -m 'Add NewFeature'`).
4. Push to the branch (`git push origin feature/NewFeature`).
5. Open a Pull Request.

## Authors
- Sayed Ahmed Husain
- Qasim Aljaffer
- Mohammed AlAlawi
- Abdulla Alasmawi

## License
This project is licensed under the MIT License. See the `LICENSE.md` file for details.