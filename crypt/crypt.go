package crypt

import (
	"log"
	"math/bits"
)

func Crypt(src uint8, secretKey uint8, decrypt bool) uint8 {
	left, right := src>>4, src&0xF

	log.Printf("FF: %04b %04b", left, right)

	for i := 0; i < 8; i++ {
		left ^= SBlock(i, secretKey, decrypt)
		left, right = right, left
	}

	return left<<4 | right
}

func SBlock(i int, secret uint8, decrypt bool) uint8 {
	key, r := key(i, secret, decrypt)

	h1 := 1 ^ (r>>1)&1 ^ (r>>3)&1&(r>>2)&1 ^ (r>>2)&1&(r>>1)&1&r&1 ^ (r>>3)&1&(r>>2)&1&(r>>1)&1&r&1 ^ (key>>3)&1
	h2 := 1 ^ (r>>3)&1&r&1 ^ (r>>3)&1&(r>>1)&1&r&1 ^ (r>>3)&1&(r>>2)&1&(r>>1)&1&r&1 ^ (key>>2)&1
	h3 := (r>>2)&1 ^ (r>>3)&1&r&1 ^ (r>>3)&1&(r>>1)&1&r&1 ^ (r>>3)&1&(r>>2)&1&(r>>1)&1&r&1 ^ (key>>1)&1
	h4 := (r>>3)&1&r&1 ^ (r>>2)&1&(r>>1)&1 ^ (r>>3)&1&(r>>2)&1&r&1 ^ (r>>3)&1&(r>>1)&1&r&1 ^ (r>>3)&1&(r>>2)&1&(r>>1)&1&r&1 ^ key&1

	return h1<<3 | h2<<2 | h3<<1 | h4
}

func key(i int, secret uint8, decrypt bool) (key uint8, rightKey uint8) {
	var (
		l uint8
		r uint8
	)

	l, r = secret>>4, secret&0xF

	if i < 7 {
		l, r = RotateLeft4(l, i), RotateLeft4(r, i)

		secret = bits.RotateLeft8(l<<4|r, 1)
	}

	return secret, r
}

func RotateLeft4(x uint8, k int) uint8 {
	const n = 4
	s := uint(k) & (n - 1)
	return (x<<s | x>>(n-s)) & 0xF
}
