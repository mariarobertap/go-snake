package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

var Snake [300][2]int
var pontos = 1
var Xsnake = 2
var Ysnake = 2
var XApple = 30
var YApple = 5

func gotoxyAndPrint(x, y int, textColor, bgColor termbox.Attribute, message string) {
	termbox.SetCursor(x, y)
	for i, char := range message {
		termbox.SetCell(x+i, y, char, textColor, bgColor)
	}
	termbox.SetCursor(0, 0)

}

func DrawInterface(width int, height int) {
	fmt.Printf("\u2554")

	i := 0
	j := 0
	for i = 0; i < width; i++ {
		fmt.Printf("\u2550")
	}

	fmt.Printf("\u2557")

	fmt.Printf("\n")

	for j = 0; j < height; j++ {

		fmt.Printf("\u2551")
		for i := 0; i < width; i++ {

			fmt.Printf(" ")

		}
		fmt.Printf("\u2551\n")
	}

	fmt.Printf("\u255A")

	for i := 0; i < width; i++ {
		fmt.Printf("\u2550")
	}

	fmt.Printf("\u255D")

	return
}

func DrawSnake() {

	if pontos == 0 {
		return
	}

	for i := 0; i < pontos; i++ {

		gotoxyAndPrint(Snake[i][0], Snake[i][1], termbox.ColorGreen, termbox.ColorDefault, "◮")

	}
	return
}

func DrawApple() {
	rand.Seed(time.Now().UnixNano())

	YApple = rand.Intn(13) + 1
	XApple = rand.Intn(28) + 1

	gotoxyAndPrint(XApple, YApple, termbox.ColorRed, termbox.ColorDefault, "ó")

}

func updateSnakeBodyAndPosition() {

	if pontos == 0 {
		return
	}

	gotoxyAndPrint(Snake[pontos][0], Snake[pontos][1], termbox.ColorWhite, termbox.ColorDefault, " ")

	for i := pontos; i >= 0; i-- {
		Snake[i+1][0] = Snake[i][0]
		Snake[i+1][1] = Snake[i][1]
	}

	return
}

func DrawScoreBoard() {

	if pontos == 0 {
		return
	}

	debugMessage := fmt.Sprintf("SCORE: [%d] HIGH SCORE: [%d] LEVEL: [%d] ", pontos, pontos, 1)

	gotoxyAndPrint(0, 17, termbox.ColorGreen, termbox.ColorDefault, debugMessage)
}

func Collide() bool {

	for i := 1; i < pontos; i++ {

		if Xsnake == Snake[i][0] && Ysnake == Snake[i][1] {
			return true
		}
	}
	return false
}

func CollideWithWalll() bool {

	if Ysnake == 1 || Ysnake == 14 {
		return true
	}
	if Xsnake == 1 || Xsnake == 29 {
		return true
	}

	return false
}

func main() {

	var keyPressed termbox.Key
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	DrawInterface(30, 15)
	DrawApple()
	for {
		Snake[0][0] = Xsnake
		Snake[0][1] = Ysnake
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
			if keyPressed == termbox.KeyArrowDown {
				if pontos != 0 {
					Ysnake++
				}
			} else if keyPressed == termbox.KeyArrowUp {
				if pontos != 0 {
					Ysnake--
				}
			} else if keyPressed == termbox.KeyArrowLeft {
				if pontos != 0 {
					Xsnake--
				}
			} else if keyPressed == termbox.KeyArrowRight {
				if pontos != 0 {
					Xsnake++
				}
			} else if keyPressed == termbox.KeyEnter {

				if pontos == 0 {
					DrawInterface(30, 15)
					DrawApple()
					pontos = 1
					Xsnake = 2
					Ysnake = 2
				}

			}

			if Xsnake == XApple && Ysnake == YApple {
				pontos++
				DrawApple()
			}

			time.Sleep(100 * time.Millisecond)

			if Collide() || CollideWithWalll() {
				//Reset snake

				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				termbox.Flush()

				//gotoxyAndPrint(0, 0, termbox.ColorGreen, termbox.ColorDefault, "Você perdeu!")
				pontos = 0

				gotoxyAndPrint(5, 10, termbox.ColorWhite, termbox.ColorDefault, "Você perdeu!")
				gotoxyAndPrint(5, 12, termbox.ColorWhite, termbox.ColorDefault, "Aperte 'ENTER' para jogar de novo")

			}
			updateSnakeBodyAndPosition()
			DrawSnake()
			DrawScoreBoard()

		}
		termbox.Flush()

	}

}
