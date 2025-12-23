package utils

import (
	"math/rand"
	"time"
)

func HumanDelay(min, max int) {
	time.Sleep(time.Duration(rand.Intn(max-min)+min) * time.Millisecond)
}
