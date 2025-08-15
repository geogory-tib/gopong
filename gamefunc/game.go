package gamefunc

import (
	"math/rand"
	"pong/types"
)

func CheckBallCollision(PlayerPaddle, CompPaddle types.Paddle, Ball *types.Ball, Mwidth, Mheight int) {
	ballCollison := false
	var y int // where the ball collision accures on the
	if Ball.X == PlayerPaddle.X {
		for y = PlayerPaddle.Y; y < PlayerPaddle.Y2; y++ {
			if Ball.Y == y {
				ballCollison = true
				break
			}
		}
		if ballCollison == true {
			Ball.DirX = 1
			if y <= PlayerPaddle.Y+1 {
				Ball.DirY = -1
			} else if y >= PlayerPaddle.Y2-1 {
				Ball.DirY = 1
			} else {
				Ball.DirY = 0
			}

		} else {
			Ball.X = Mwidth / 2
			Ball.Y = Mheight / 2
			Ball.DirX = 1
			Ball.DirY = 0
		}
	} else if Ball.X == CompPaddle.X {
		for y = CompPaddle.Y; y < CompPaddle.Y2; y++ {
			if Ball.Y == y {
				ballCollison = true
				break
			}
		}
		if ballCollison == true {
			Ball.DirX = -1
			if y <= CompPaddle.Y+1 {
				Ball.DirY = -1
			} else if y >= CompPaddle.Y2-1 {
				Ball.DirY = 1
			} else {
				Ball.DirY = 0
			}

		} else {
			Ball.X = Mwidth / 2
			Ball.Y = Mheight / 2
			Ball.DirX = -1
			Ball.DirY = 0
		}
	} else if Ball.Y == 0 {
		Ball.DirY = 1
	} else if Ball.Y == Mheight {
		Ball.DirY = -1
	}
}

func CompPaddleAi(Ball types.Ball, Paddle *types.Paddle, Mwidth, Mheight int) {
	AiMiss := rand.Intn(15) // random chance the ai does nothing and potentially misses the ball

	if AiMiss == 4 || Ball.DirX == -1 {
		return
	}
	if Paddle.Y+3 >= Ball.Y {
		Paddle.Y--
		Paddle.Y2--
	}
	if Paddle.Y+3 <= Ball.Y {
		Paddle.Y++
		Paddle.Y2++
	}
}
