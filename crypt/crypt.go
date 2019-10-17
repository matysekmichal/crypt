package crypt

import (
	"math"
	"math/bits"
)

func Crypt(msg uint8, secretKey uint8, decrypt bool) uint8 {
	left, right := msg>>4, msg&0xF

	for i := 1; i <= 8; i++ {
		left ^= SBlock(i, right, secretKey, decrypt)

		if i < 8 {
			left, right = right, left
		}
	}

	return left<<4 | right
}

func SBlock(iteration int, r uint8, secretKey uint8, decrypt bool) uint8 {
	key := RoundKeyGenerator(iteration, secretKey, decrypt)

	h1 := 1 ^ (r>>1)&1 ^ (r>>3)&1&(r>>2)&1 ^ (r>>2)&1&(r>>1)&1&r&1 ^ (r>>3)&1&(r>>2)&1&(r>>1)&1&r&1 ^ (key>>7)&1
	h2 := 1 ^ (r>>3)&1&r&1 ^ (r>>3)&1&(r>>1)&1&r&1 ^ (r>>3)&1&(r>>2)&1&(r>>1)&1&r&1 ^ (key>>5)&1
	h3 := (r>>2)&1 ^ (r>>3)&1&r&1 ^ (r>>3)&1&(r>>1)&1&r&1 ^ (r>>3)&1&(r>>2)&1&(r>>1)&1&r&1 ^ (key>>3)&1
	h4 := (r>>3)&1&r&1 ^ (r>>2)&1&(r>>1)&1 ^ (r>>3)&1&(r>>2)&1&r&1 ^ (r>>3)&1&(r>>1)&1&r&1 ^ (r>>3)&1&(r>>2)&1&(r>>1)&1&r&1 ^ (key>>1)&1

	return h1<<3 | h2<<2 | h3<<1 | h4
}

func RoundKeyGenerator(iteration int, secretKey uint8, decrypt bool) uint8 {
	var (
		left  uint8
		right uint8
	)

	if decrypt {
		iteration = int(math.Abs(float64(iteration - 9)))
	}

	for i := 0; i < iteration; i++ {
		left, right = secretKey>>4, secretKey&0xF

		if i%2 == 0 {
			left, right = RotateLeft4(left, 1), RotateLeft4(right, 1)
			secretKey = left<<4 | right
		} else {
			secretKey = bits.RotateLeft8(secretKey, 1)
		}
	}

	return secretKey
}

func RotateLeft4(x uint8, k int) uint8 {
	const n = 4
	s := uint(k) & (n - 1)
	return (x<<s | x>>(n-s)) & 0xF
}
