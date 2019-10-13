package main

import (
	"fmt"
	"github.com/matysekmichal/mpc.crypt/crypt"
)

func main() {
	var (
		operation = "S"
		str string
		decrypt = false
	)

	fmt.Println("Program pozwalający na szyfrowanie i deszyfrowanie bloków 8-bitowych.")
	fmt.Print("Chcesz (S)zyfrować, czy (D)eszyfrować (DOMYŚLNIE S): ")
	var _, err = fmt.Scan(&operation)

	if err != nil || operation != "S" && operation != "D" {
		fmt.Println(err)
		panic("Wprowadzono nieprawidłową wartość.")
	}

	if operation == "D" {
		decrypt = true
	}

	fmt.Print("Wprowadź klucz 8-bitowy: ")
	_, err = fmt.Scan(&str)

	if err != nil || len(str) != 8 {
		panic("Wprowadzony klucz jest nieprawidłowy.")
	}



	result := crypt.Crypt(str, decrypt)

	fmt.Printf("%b", result)
}
