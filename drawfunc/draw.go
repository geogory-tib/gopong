package drawfunc

import (
	"pong/types"

	"github.com/nsf/termbox-go"
)

func DrawScreen(compPaddle, playerPaddle types.Paddle, ball types.Ball, Mwidth, Mheight int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y := playerPaddle.Y; y < playerPaddle.Y2; y++ {
		termbox.SetBg(playerPaddle.X, y, termbox.ColorWhite)
	}
	for y := compPaddle.Y; y < compPaddle.Y2; y++ {
		termbox.SetBg(compPaddle.X, y, termbox.ColorWhite)
	}
	termbox.SetChar(ball.X, ball.Y, 'o')
	termbox.Flush()
}
