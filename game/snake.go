package game

type snake struct {
	length int
	Head *coord
	Tail []*coord
	tailLength int
}

func (s *snake) getLastTail() (int, int) {
	return s.Tail[len(s.Tail) - 1].get()
}

func (s *snake) addTail(dir direction)  {
	xH, yH := s.Head.get()
	newTail := &coord{
		x: xH,
		y: yH,
	}
	if len(s.Tail) > 0 {
		xH, yH = s.Tail[len(s.Tail) - 1].get()
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
	newTail.set(xH, yH)
	s.Tail = append(s.Tail, newTail)
}

type coord struct {
	x int
	y int
}

func (c *coord) set(x,y int)  {
	c.x = x
	c.y = y
}

func (c *coord) get() (int, int) {
	return c.x, c.y
}
