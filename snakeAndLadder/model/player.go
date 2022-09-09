package model

type Player struct {
	name string
	pos  int
}

func NewPlayer(name string) *Player {
	return &Player{
		name: name,
		pos:  0,
	}
}

func GenerateAllPlayers(names []string, n int) []*Player {
	res := make([]*Player, n)

	for i, name := range names {
		res[i] = NewPlayer(name)
	}
	return res
}

func (p *Player) UpdatePos(newPos int) {
	p.pos = newPos
}

func (p *Player) GetPos() int {
	return p.pos
}

func (p *Player) GetName() string {
	return p.name
}
