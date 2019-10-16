package crypt

import (
	"math"
	"math/bits"
)

func Crypt(msg uint8, secretKey uint8, decrypt bool) uint8 {
	left, right := msg>>4, msg&0xF

	for i := 0; i < 8; i++ {
		left ^= SBlock(i, right, secretKey, decrypt)

		if i < 7 {
			left, right = right, left
		}
	}

	return left<<4 | right
}

func SBlock(i int, r uint8, secret uint8, decrypt bool) uint8 {
	key := RoundKeyGenerator(i, secret, decrypt)

	h1 := 1 ^ (r>>1)&1 ^ (r>>3)&1&(r>>2)&1 ^ (r>>2)&1&(r>>1)&1&r&1 ^ (r>>3)&1&(r>>2)&1&(r>>1)&1&r&1 ^ (key>>3)&1
	h2 := 1 ^ (r>>3)&1&r&1 ^ (r>>3)&1&(r>>1)&1&r&1 ^ (r>>3)&1&(r>>2)&1&(r>>1)&1&r&1 ^ (key>>2)&1
	h3 := (r>>2)&1 ^ (r>>3)&1&r&1 ^ (r>>3)&1&(r>>1)&1&r&1 ^ (r>>3)&1&(r>>2)&1&(r>>1)&1&r&1 ^ (key>>1)&1
	h4 := (r>>3)&1&r&1 ^ (r>>2)&1&(r>>1)&1 ^ (r>>3)&1&(r>>2)&1&r&1 ^ (r>>3)&1&(r>>1)&1&r&1 ^ (r>>3)&1&(r>>2)&1&(r>>1)&1&r&1 ^ key&1

	return h1<<3 | h2<<2 | h3<<1 | h4
}

func RoundKeyGenerator(i int, secret uint8, decrypt bool) uint8 {
	var l, r = secret >> 4, secret & 0xF
	move := 1

	if decrypt {
		i = int(math.Abs(float64(i - 7)))
	}

	if i > 6 {
		i -= 1
		move += 1
	}

	l, r = RotateLeft4(l, i), RotateLeft4(r, i)
	secret = bits.RotateLeft8(l<<3|r, move)

	return secret
}

func RotateLeft4(x uint8, k int) uint8 {
	const n = 4
	s := uint(k) & (n - 1)
	return (x<<s | x>>(n-s)) & 0xF
}
