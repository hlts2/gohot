package gohot

import "testing"

func BenchmarkTestCreateOneHotVectorFromTokens(b *testing.B) {
	b.StopTimer()

	for i := 0; i < b.N; i++ {
		tokens := []string{"私", "は", "ブドウ", "が", "好き", "な", "ので", "、", "毎年", "秋", "が", "楽しみ", "だ"}
		b.StartTimer()

		_ = CreateOneHotVectorFromTokens(tokens)

		b.StartTimer()
	}
}

func BenchmarkTestCreateOneHotVectorFromText(b *testing.B) {
	b.StopTimer()

	for i := 0; i < b.N; i++ {
		str := "私はブドウが好きなので、毎年秋が楽しみだ"
		b.StartTimer()

		_ = CreateOneHotVectorFromText(str)

		b.StartTimer()
	}
}
