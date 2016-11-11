package util

import (
	"math/rand"
	"time"
)

func GetCurrentSecond() int64 {
	return time.Now().Unix()
}

func gen(index int, r *rand.Rand) int {
	return r.Intn(index)
}

func RandShuffle(arr []int, first int, last int) {
	if first >= last {
		return
	}
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	n := last - first
	for i := n - 1; i > 0; i-- {
		index := gen(i+1, r)
		arr[i], arr[index] = arr[index], arr[i]
	}
}
