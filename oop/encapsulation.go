// encapsulation is the process of keeping data private and exposing only what is necessary to external sources

// In Go, the first letter of a variable designates if it is public or private.
// Lowercase for unexported and capital letters for exported variables.
// variables with lowercase names are invisible to other packages in Go, while uppercase vars can be accessed by other code that import this package

package main

type Player struct {
	Team     string
	Name     string
	Position string
}

func (p *Player) GetTeam() string {
	return p.Team
}

func (p *Player) GetName() string {
	return p.Name
}

func (p *Player) GetPosition() string {
	return p.Position
}
