package utils

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func HumanType(el *rod.Element, text string) {
	for _, ch := range text {
		el.MustInput(string(ch))
		time.Sleep(time.Duration(rand.Intn(120)+80) * time.Millisecond)
	}
}
