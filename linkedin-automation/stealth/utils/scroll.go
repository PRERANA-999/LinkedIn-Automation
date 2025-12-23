package utils

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func RandomScroll(page *rod.Page) {
	for i := 0; i < 3; i++ {
		page.Mouse.Scroll(0, rand.Intn(500)+200)
		time.Sleep(time.Duration(rand.Intn(800)+400) * time.Millisecond)
	}
}
