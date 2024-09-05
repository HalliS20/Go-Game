package player

type Player struct {
	Name     string
	x        int
	y        int
	rotation int
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:     name,
		x:        0,
		y:        0,
		rotation: 0,
	}
}

func (p *Player) GetPosition() (int, int) {
	return p.x, p.y
}

// TODO: May Be improved
func (p *Player) MoveForward() {
	if p.GetRotation() == 0 {
		p.y++
	} else if p.GetRotation() == 90 {
		p.x++
	} else if p.GetRotation() == 180 {
		p.y--
	} else if p.GetRotation() == 270 {
		p.x--
	}
}

func (p *Player) GetRotation() int {
	return p.rotation % 360
}

func (p *Player) RotateRight() {
	p.rotation += 90
}
