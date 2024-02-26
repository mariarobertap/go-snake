package game

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type Apple struct {
	XPosition int
	YPosition int
}

func NewApple() *Apple {

	//TODO: Is this okay
	return &Apple{
		XPosition: 30,
		YPosition: 5,
	}
}

func (a *Apple) Draw() {
	rand.Seed(time.Now().UnixNano())

	//Turn this into an apple function
	a.YPosition = rand.Intn(14) + 1
	a.XPosition = rand.Intn(39) + 1

	gotoxyAndPrint(a.XPosition, a.YPosition, termbox.ColorRed, termbox.ColorDefault, "ó")

}
