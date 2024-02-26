package game

import "github.com/nsf/termbox-go"

type Snake struct {
	BodyVectorPosition [300][2]int
	XPosition          int
	YPosition          int
	SnakeSize          int
}

func NewSnake() *Snake {

	//TODO: Is this okay
	return &Snake{
		XPosition: 2,
		YPosition: 2,
		SnakeSize: 1,
	}
}

func (s *Snake) drawSnake() {

	if s.SnakeSize == 0 {
		return
	}

	for i := 0; i < s.SnakeSize; i++ {

		gotoxyAndPrint(s.BodyVectorPosition[i][0], s.BodyVectorPosition[i][1], termbox.ColorGreen, termbox.ColorDefault, "◮")

	}
	return
}

func (s *Snake) selfCollided() bool {
	for i := 1; i < s.SnakeSize; i++ {

		if s.XPosition == s.BodyVectorPosition[i][0] && s.YPosition == s.BodyVectorPosition[i][1] {
			return true
		}
	}
	return false
}

func (s *Snake) collideWithWall(InterfaceHeigth int, InterfaceWidth int) bool {

	if s.YPosition == -1 || s.YPosition == InterfaceHeigth+1 {
		return true
	}
	if s.XPosition == -1 || s.XPosition == InterfaceWidth+1 {
		return true
	}

	return false
}

func (s *Snake) updateSnakeBody() {

	if s.SnakeSize == 0 {
		return
	}

	gotoxyAndPrint(s.BodyVectorPosition[s.SnakeSize][0], s.BodyVectorPosition[s.SnakeSize][1], termbox.ColorWhite, termbox.ColorDefault, " ")

	for i := s.SnakeSize; i >= 0; i-- {
		s.BodyVectorPosition[i+1][0] = s.BodyVectorPosition[i][0]
		s.BodyVectorPosition[i+1][1] = s.BodyVectorPosition[i][1]
	}

	return
}

func (s *Snake) updateSnakeHead() {
	s.BodyVectorPosition[0][0] = s.XPosition
	s.BodyVectorPosition[0][1] = s.YPosition
}
