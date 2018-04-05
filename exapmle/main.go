package main

import (
	"fmt"

	"github.com/hlts2/gohot"
)

func main() {
	onehotvectors := gohot.CreateOneHotVectorFromTokens([]string{"私", "は", "ブドウ", "が", "好き", "な", "ので", "、", "毎年", "秋", "が", "楽しみ", "だ"})

	for token, vector := range onehotvectors {
		fmt.Println(token, vector)
	}

	onehotvectors = gohot.CreateOneHotVectorFromText("私はブドウが好きなので、毎年秋が楽しみだ")

	for token, vector := range onehotvectors {
		fmt.Println(token, vector)
	}
}
