package random

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//StringFromSet ...
func StringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

//Bool ...
func Bool() bool {
	return rand.Intn(2) == 1
}

//Int64 ...
func Int(min, max int64) int64 {
	return min + rand.Int63()%(max-min+1)
}

//Float64 ...
func Float64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// Float32 ..
func Float32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

//ID efrwefg
func ID() string {
	return uuid.New().String()
}
