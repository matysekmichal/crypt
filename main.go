package main

import (
	"fmt"
	"github.com/matysekmichal/mpc.crypt/crypt"
	"strconv"
)

var (
	operation = "S"
	decrypt   = false
	str       string
	key       string
)

func main() {
	dataProvider()

	msg, _ := strconv.ParseInt(str, 2, 16)
	secretKey, _ := strconv.ParseInt(key, 2, 16)

	a := crypt.Crypt(uint8(msg), uint8(secretKey), decrypt)

	fmt.Printf("WYNIK: %08b (%02X)", a, a)

	_, _ = fmt.Scanln()
}

func dataProvider() {
	fmt.Println("Program pozwalający na szyfrowanie i deszyfrowanie bloków 8-bitowych.")
	fmt.Print("Chcesz (S)zyfrować, czy (D)eszyfrować (DOMYŚLNIE S): ")
	var _, err = fmt.Scan(&operation)

	if err != nil || operation != "S" && operation != "D" {
		panic("Wprowadzono nieprawidłową wartość.")
	}

	if operation == "D" {
		decrypt = true
	}

	fmt.Print("Wprowadź wiadomość 8-bitową: ")
	_, err = fmt.Scan(&str)

	if err != nil || len(str) != 8 {
		panic("Wprowadzony wiadomość jest nieprawidłowa.")
	}

	fmt.Print("Wprowadź klucz szyfrujący 8-bitowy: ")
	_, err = fmt.Scan(&key)

	if err != nil || len(key) != 8 {
		panic("Wprowadzony klucz szyfrujący jest nieprawidłowy.")
	}
}
