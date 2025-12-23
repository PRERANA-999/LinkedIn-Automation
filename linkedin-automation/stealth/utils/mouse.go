package utils

import (
	"time"

	"github.com/go-rod/rod"
)

func HumanClick(page *rod.Page, el *rod.Element) {
	box := el.MustShape().Box()
	page.Mouse.Move(box.X+10, box.Y+10, 5)
	time.Sleep(300 * time.Millisecond)
	el.MustHover()
}
