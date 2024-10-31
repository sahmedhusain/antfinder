package structs

// State represents the overall state of the system
type State struct {
	Visited         map[string]bool     // Tracks visited rooms
	Tunnels         map[string][]string // Maps rooms to their connected tunnels
	Rooms           []string            // List of all rooms
	RoomsMapRoom    map[string]bool     // Maps room names to a boolean indicating their presence
	RoomsMapTunnels map[string]string   // Maps room names to their corresponding tunnels
	StartEnd        map[string]int      // Maps start and end room names to their indices
	StartEndRooms   map[string]string   // Maps start and end room names to their corresponding room names
	Ants            int                 // Number of ants
	AntsHere        bool                // Indicates if ants are present
	Solution        map[string]string   // Maps room names to their solution paths
	RoomCoordinates map[string]int      // Maps room names to their coordinates
	Paths           [][]string          // List of paths for the ants
	End             string              // Name of the end room
	StartAndEnd     []string            // List containing start and end room names
	Connect         map[string]int      // Maps room names to their connection counts
	Many            map[string]int      // Maps room names to some integer value (purpose not specified)
}
