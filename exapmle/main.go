package main

import (
	"fmt"

	"github.com/hlts2/gohot"
)

func main() {
	vectors := gohot.CreateOneHotVectorFromTokens([]string{"私", "は", "ブドウ", "が", "好き", "な", "ので", "、", "毎年", "秋", "が", "楽しみ", "だ"})

	for key, val := range vectors {
		fmt.Println(key, val)
	}

	vectors = gohot.CreateOneHotVectorFromText("私はブドウが好きなので、毎年秋が楽しみだ")

	for key, val := range vectors {
		fmt.Println(key, val)
	}
}
