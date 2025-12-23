package utils

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func HumanScroll(page *rod.Page, times int) {
	for i := 0; i < times; i++ {
		page.MustEval(`() => {
			window.scrollBy({
				top: Math.floor(Math.random() * 400) + 200,
				behavior: "smooth"
			});
		}`)
		time.Sleep(time.Duration(rand.Intn(1200)+800) * time.Millisecond)
	}
}
