package util

import (
	"bytes"
	"math/rand"
	"time"
)

var r *rand.Rand

const chars = "abcdefghijklmnopqrstuvwxyz"

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandUrl() string {
	list := []string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.twitter.com",
		"https://www.facebook.com",
	}

	index := r.Intn(len(list) - 1)
	return list[index]
}

func RandInt(min, max int64) int64 {
	return min + r.Int63n(max-min+1)
}

func RandString(n int) string {
	var buffer bytes.Buffer
	l := len(chars)

	for i := 0; i < n; i++{
		buffer.WriteByte(chars[r.Intn(l)])
	}

	return buffer.String()
}
