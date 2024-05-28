package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	numerator, denominator := 20, 3
	remainder, mod, err := calcRemainderAndMod(numerator, denominator)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%d÷%d = %d余り%d\n", numerator, denominator, remainder, mod)
}

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("0で割れません")
	}
	remainder := numerator / denominator
	mod := numerator % denominator
	return remainder, mod, nil
}
