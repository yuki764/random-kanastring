package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage:", os.Args[0], "<length>")
		os.Exit(1)
	}
	const KANA = "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをんアイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲン"

	length, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len([]rune(KANA)))))
		if err != nil {
			panic(err)
		}
		moji := string([]rune(KANA)[n.Int64()])
		fmt.Print(moji)
	}
	fmt.Println()
}
