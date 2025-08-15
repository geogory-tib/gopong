package types

const DIR_UP = 1
const DIR_DOWN = -1
const DIR_SAME = 0

type Paddle struct {
	X, Y, Y2 int
	Dir      int
}

type Ball struct {
	X, Y       int
	DirX, DirY int
}

func (self *Paddle) Init(player bool, Mwidth, Mheight int) {
	if player == true {
		self.X = 0
	} else {
		self.X = Mwidth - 1
	}
	self.Y = (Mheight - 5) / 2
	self.Y2 = self.Y + 5
	self.Dir = DIR_SAME

}
func (self *Ball) Init(Mwidth, Mheight int) {
	self.X = Mwidth / 2
	self.Y = Mheight / 2
	self.DirX = -1
	self.DirY = 0
}
func (self *Ball) UpdateBallPos() {
	self.X += self.DirX
	self.Y += self.DirY
}
