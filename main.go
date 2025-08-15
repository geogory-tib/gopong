package main

import (
	"log"
	"pong/drawfunc"
	"pong/gamefunc"
	"pong/types"
	"pong/ui"
	"time"

	"github.com/nsf/termbox-go"
)

func eventPoll(playerPaddle *types.Paddle) {
	for {
		event := termbox.PollEvent()
		if event.Key == termbox.KeyArrowDown {
			playerPaddle.Y++
			playerPaddle.Y2++
		}
		if event.Key == termbox.KeyArrowUp {
			playerPaddle.Y--
			playerPaddle.Y2--
		}
		if event.Key == termbox.KeyEsc {
			playerPaddle.Y = -20
		}
	}
}
func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()
	Mwidth, Mheight := termbox.Size()
	if ui.Menu(Mwidth, Mheight) == 1 {
		return
	}
	//var frameCounter int //used for the ai
	var ball types.Ball
	var playerPaddle, compPaddle types.Paddle
	ball.Init(Mwidth, Mheight)
	playerPaddle.Init(true, Mwidth, Mheight)
	compPaddle.Init(false, Mwidth, Mheight)
	drawfunc.DrawScreen(compPaddle, playerPaddle, ball, Mheight, Mwidth)
	go eventPoll(&playerPaddle)
	for { // game loop
		drawfunc.DrawScreen(compPaddle, playerPaddle, ball, Mheight, Mwidth)
		ball.UpdateBallPos()
		gamefunc.CheckBallCollision(playerPaddle, compPaddle, &ball, Mwidth, Mheight)
		gamefunc.CompPaddleAi(ball, &compPaddle, Mwidth, Mheight)
		if playerPaddle.Y == -20 {
			return
		}
		time.Sleep(50 * time.Millisecond)

	}
}
