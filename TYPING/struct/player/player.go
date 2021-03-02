package player

// Avatar ...
type Avatar struct {
	Url string
}

// Player ...
type Player struct {
	Name     string
	Age      int
	Avatar   Avatar
	password string
}

// New returns Player
func New(name string) Player {
	return Player{
		Name:     name,
		password: "defaultpassword",
	}
}
