package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var signs = []byte("QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm")
	res := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range res {
		res[i] = signs[rand.Intn(len(signs))]
	}
	return string(res)
}
