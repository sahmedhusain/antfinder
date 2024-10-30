package structs

type State struct {
	Visited         map[string]bool
	Tunnels         map[string][]string
	Rooms           []string
	RoomsMapRoom    map[string]bool
	RoomsMapTunnels map[string]string
	StartEnd        map[string]int
	StartEndRooms   map[string]string
	Ants            int
	AntsHere        bool
	Solution        map[string]string
	RoomCoordinates map[string]int
	Paths           [][]string
	End             string
	StartAndEnd     []string
	Connect         map[string]int
	Many            map[string]int
}
