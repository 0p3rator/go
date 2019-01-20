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

func newFood(e rune, x int, y int) *food {
	return &food{
		points: 10,
		emoji: getEmoji(e),
		x: x,
		y: y,
	}
}

func getEmoji(e rune) rune {
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