package snake

import (
	"time"

	"github.com/nsf/termbox-go"
)

var (
	pointsChan =  make(chan int)
	keyboardEventsChan =  make(chan keyboardEvent)
)
type game struct {
	arena *arena
	score int
	isOver bool
}

func initSnake() *snake {
	return newSnake(RIGHT,[]coord{
		coord{x: 1, y: 1},
		coord{x: 1, y: 2},
		coord{x: 1, y: 3},
		coord{x: 1, y: 4},
	})
}

func initArena() *arena {
	return newArena(initSnake(), pointsChan, 20, 50)
}

func initScore() int {
	return 0
}


func NewGame() *game {
	return &game{
		arena: initArena(),
		score: initScore(),
	}
}

func (g *game) moveInterval() time.Duration {
	ms := 100 - (g.score / 10)
	return time.Duration(ms) * time.Millisecond
}

func (g *game) end() {
	g.isOver = true
}

func (g *game) retry() {
	g.arena = initArena()
	g.score = initScore()
	g.isOver = false
}

func (g *game) addPoints(p int) {
	g.score += p
}
func (g *game) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	defer termbox.Close()

	go listenToKeyboardEvent(keyboardEventsChan)

	if err := g.render(); err != nil {
		panic(err)
	}

mainloop:
	for {
		select {
		case p := <- pointsChan:
			g.addPoints(p)
		case e := <- keyboardEventsChan:
			switch e.eventType {
			case MOVE:
				d := keyToDirection(e.key)
				g.arena.snake.changeDirection(d)
			case RETRY:
				g.retry()
			case END:
				break mainloop
			}
		default:
			if !g.isOver {
				if err := g.arena.moveSnake(); err != nil {
					g.end()
				}
			}

			if err := g.render(); err != nil {
				panic(err)
			}
			
			time.Sleep(g.moveInterval())
		}
	}

}

