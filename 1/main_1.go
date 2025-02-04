package main

import (
	"1/pkg"
	"fmt"
)

var (
	numDecimal               = 42
	numOctal                 = 052
	numHexadecimal           = 0x2A
	pi                       = 3.14
	name                     = "Golang"
	isActive                 = true
	complexNum     complex64 = 1 + 2i
)

func main() {
	vars := []interface{}{numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum}
	for _, v := range vars {
		fmt.Printf("Значение переменной: %v, тип: %T\n", v, v)
	}

	combined := pkg.StrConv(vars...)
	fmt.Printf("Combined string %s\n", combined)

	hashedResult := pkg.AddSaltAndHash(combined)
	fmt.Printf("Hashed result: %s\n", hashedResult)
}
