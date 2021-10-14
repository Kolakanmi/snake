package game

type Snake struct {
	length int
	Head *coord
	Tail []*coord
	tailLength int
}

func CreateSnake(x, y int) *Snake {
	return &Snake{
		length: 1,
		Head: &coord{
			x: x,
			y: y,
		},
	}
}

func (s *Snake) GetLastTail() (int, int) {
	return s.Tail[len(s.Tail) - 1].Get()
}

func (s *Snake) AddTail(dir Direction)  {
	xH, yH := s.Head.Get()
	newTail := &coord{
		x: xH,
		y: yH,
	}
	if len(s.Tail) > 0 {
		xH, yH = s.Tail[len(s.Tail) - 1].Get()
		//newTail.Set(xH, yH)
	}
	switch dir {
	case left:
		xH++
	case right:
		xH--
	case up:
		yH++
	case down:
		yH--
	}
	newTail.Set(xH, yH)
	s.Tail = append(s.Tail, newTail)
}

type coord struct {
	x int
	y int
}

func (c *coord) Set(x,y int)  {
	c.x = x
	c.y = y
}

func (c *coord) Get() (int, int) {
	return c.x, c.y
}