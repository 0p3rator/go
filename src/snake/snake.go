package snake

import "errors"

// direction int use to indicate snake's move direction
type direction int

// direction have four default values, init with increase number from 1
const (
	RIGHT direction = 1 + iota
	LEFT
	UP
	DOWN
)

// struct snake use to store snake's info
type snake struct {
	direction direction
	body []coord
	length int
}

// snake's custom constructor with args direction and body
func newSnake(d direction, b []coord) *snake {
	return &snake{
		direction: d,
		body: b,
		length: len(b),
	}
}

// snake's head use to represent snake
func (s *snake) head() coord {
	return s.body[len(s.body) - 1]
}

// snake die with throw "Died" error
func (s *snake) die() error {
	return errors.New("Died")
}

// snake change move direction after keyboard event, if direction is opposite to snake's current direction, do nothing.
func (s *snake) changeDirection(d direction) {
	opposites := map[direction]direction {
		RIGHT: LEFT,
		LEFT: RIGHT,
		UP: DOWN,
		DOWN: UP,
	}
	if o := opposites[d]; o != 0 && o != s.direction {
		s.direction = d
	}
}

// snake move to specific direction; snake dies if move to its body; snake's body increase 1 if snake's length increases when snake eat food
func (s *snake) move() error {
	h := s.head()
	c := coord{x: h.x, y: h.y}
	switch s.direction {
	case RIGHT:
		c.x++
	case LEFT:
		c.x--
	case UP:
		c.y++
	case DOWN:
		c.y--
	}
	if s.isOnPosition(c) {
		return s.die()
	}

	if s.length == len(s.body) {
		s.body = append(s.body[1:], c)
	} else if s.length > len(s.body) {
		s.body = append(s.body, c)
	}
	
	return nil
}

// snake move to his own body if head's next step is in his body coord list
func (s *snake) isOnPosition(c coord) bool {
	for _, b  := range s.body {
		if b.x == c.x && b.y == c.y {
			return true
		} 
	}
	return false
}
