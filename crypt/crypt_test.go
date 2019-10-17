package crypt

import (
	"strconv"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	dataProvider := []struct {
		key     string
		msg     string
		decrypt string
	}{
		{"11100000", "00000000", "01001111"},
		{"10110001", "01111000", "11101111"},
		{"10101100", "11111010", "11111011"},
	}

	for _, data := range dataProvider {
		msg, _ := strconv.ParseInt(data.msg, 2, 16)
		decrypt, _ := strconv.ParseInt(data.decrypt, 2, 16)
		secretKey, _ := strconv.ParseInt(data.key, 2, 16)

		encryptResult := Crypt(uint8(msg), uint8(secretKey), false)
		decryptResult := Crypt(uint8(decrypt), uint8(secretKey), true)

		if encryptResult != uint8(decrypt) {
			t.Errorf("ENCRYPT: It should be %08b but got %08b", encryptResult, uint8(decrypt))
		}

		if uint8(msg) != decryptResult {
			t.Errorf("DECRYPT: It should be %08b but got %08b", uint8(msg), decryptResult)
		}
	}
}
