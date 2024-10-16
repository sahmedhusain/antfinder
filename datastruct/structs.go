package structs

type Graph struct {
	Rooms map[string]*Room
	Start *Room
	End   *Room
}

type Room struct {
	Name        string
	Connections map[string]*Room
}
type Ant struct {
	ID   int
	Path []*Room
}


