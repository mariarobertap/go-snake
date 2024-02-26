package game

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

type Game struct {
	pontos          int
	Snake           *Snake
	Apple           *Apple
	InterfaceWidth  int
	InterfaceHeigth int
	GameOver        bool
	SpeedMsLevel    int
}

func NewGame() *Game {

	return &Game{
		pontos:          1,
		Snake:           NewSnake(),
		Apple:           NewApple(),
		InterfaceWidth:  40,
		InterfaceHeigth: 15,
		SpeedMsLevel:    100,
	}
}

func (g *Game) Run() {
	var keyPressed termbox.Key

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	//Turn this into a function
	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	g.drawInterface()
	g.Apple.Draw()

	for {
		g.Snake.updateSnakeHead()
		termbox.SetCursor(0, 0)

		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				keyPressed = ev.Key
				if ev.Key == termbox.KeyEsc {
					os.Exit(1)
				}
			}

		default:
			g.performKeyboardAction(keyPressed)

			g.checkIfSnakeAteApple()
			g.Snake.updateSnakeBody()
			g.Snake.drawSnake()
			g.drawScoreBoard()
			g.checkIfGameOver()

		}
		termbox.Flush()

		if g.GameOver {
			break
		}

	}
}

func (g *Game) drawInterface() {
	fmt.Printf("\u2554")

	i := 0
	j := 0
	for i = 0; i < g.InterfaceWidth; i++ {
		fmt.Printf("\u2550")
	}

	fmt.Printf("\u2557")

	fmt.Printf("\n")

	for j = 0; j < g.InterfaceHeigth; j++ {

		fmt.Printf("\u2551")
		for i := 0; i < g.InterfaceWidth; i++ {

			fmt.Printf(" ")

		}
		fmt.Printf("\u2551\n")
	}

	fmt.Printf("\u255A")

	for i := 0; i < g.InterfaceWidth; i++ {
		fmt.Printf("\u2550")
	}

	fmt.Printf("\u255D")

	return
}

func (g *Game) drawScoreBoard() {

	if g.pontos == 0 {
		return
	}
	//debugMessage := fmt.Sprintf("X: [%d] Y: [%d]  ", g.Snake.XPosition, g.Snake.YPosition)
	debugMessage := fmt.Sprintf("SCORE: [%d] LEVEL: [%d] ", g.pontos, 1)

	gotoxyAndPrint(0, 17, termbox.ColorGreen, termbox.ColorDefault, debugMessage)
}

func (g *Game) checkIfSnakeAteApple() {
	if g.Snake.XPosition == g.Apple.XPosition && g.Snake.YPosition == g.Apple.YPosition {
		g.pontos++
		g.Snake.SnakeSize++

		g.Apple.Draw()
	}
	time.Sleep(time.Duration(g.SpeedMsLevel) * time.Millisecond)
}

func (g *Game) checkIfGameOver() bool {

	if g.Snake.selfCollided() || g.Snake.collideWithWall(g.InterfaceHeigth, g.InterfaceWidth) {

		g.pontos = 0
		g.Snake.SnakeSize = 0
		g.GameOver = true

	}

	return true
}

func (g *Game) performKeyboardAction(keyPressed termbox.Key) {
	if keyPressed == termbox.KeyArrowDown {
		if g.pontos != 0 {
			g.Snake.YPosition++
		}
	} else if keyPressed == termbox.KeyArrowUp {
		if g.pontos != 0 {
			g.Snake.YPosition--
		}
	} else if keyPressed == termbox.KeyArrowLeft {
		if g.pontos != 0 {
			g.Snake.XPosition--
		}
	} else if keyPressed == termbox.KeyArrowRight {
		if g.pontos != 0 {
			g.Snake.XPosition++
		}
	}
}

func gotoxyAndPrint(x, y int, textColor, bgColor termbox.Attribute, message string) {
	termbox.SetCursor(x, y)
	for i, char := range message {
		termbox.SetCell(x+i, y, char, textColor, bgColor)
	}
	termbox.SetCursor(0, 0)
}
