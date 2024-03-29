package snake

import (
	"math/rand"
	"os"
	"strings"
)

type food struct {
	emoji rune
	points, x, y int	
}

func newFood(x int, y int) *food {
	return &food{
		points: 10,
		emoji: getEmoji(),
		x: x,
		y: y,
	}
}

func getEmoji() rune {
	if hasUnicodeSupport() {
		return randomEmoji()
	}
	return '@'
}

func randomEmoji() rune {
	f := []rune{
		'🍒',
		'🍍',
		'🍑',
		'🍇',
		'🍏',
		'🍌',
		'🍫',
		'🍭',
		'🍕',
		'🍩',
		'🍗',
		'🍖',
		'🍬',
		'🍤',
		'🍪',
	}
	return f[rand.Intn(len(f))]
}
func hasUnicodeSupport() bool {
	return strings.Contains(os.Getenv("LANG"), "UTF-8")
}