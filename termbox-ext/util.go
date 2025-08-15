package termboxext

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func TbPrint(width, height int, str string, fg, bg termbox.Attribute) {

	x := width
	y := height
	for _, ch := range str {
		termbox.SetCell(x, y, ch, fg, bg)
		x += runewidth.RuneWidth(ch)
	}
}

func GetStringAtLine(lineY, length int) (str string) {
	x := 0
	var retString string
	for x < length {

		cell := termbox.GetCell(x, lineY)
		retString += string(cell.Ch)
		x += runewidth.RuneWidth(cell.Ch)
	}
	retString += "\n"
	return retString
}
