package ui

import (
	termboxext "pong/termbox-ext"

	"github.com/nsf/termbox-go"
)

func Menu(Mwidth, MHeight int) uint8 {
	termboxext.TbPrint((Mwidth-len("Welcome to pong!"))/2, (MHeight/2)-3, "Welcome To Pong!", termbox.ColorBlack, termbox.ColorMagenta)
	termboxext.TbPrint((Mwidth-len("Press Enter To play or any other key to escape!"))/2, MHeight/2, "Press Enter To Play Or Any Other Key To Escape!", termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
	event := termbox.PollEvent()
	if event.Key != termbox.KeyEnter {
		return 1
	}
	return 0
}
