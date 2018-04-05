# gohot
gohot is a golang library that generates one hot vector.

## Requirement
Go 1.8

## Installation
```go
go get github.com/hlts2/gohot

```
## Example

```go
    tokens := []string{"私", "は", "ブドウ", "が", "好き", "な", "ので", "、", "毎年", "秋", "が", "楽しみ", "だ"}
    onehotvectors := gohot.CreateOneHotVectorFromTokens(tokens) //["私":"100000000000", "は": "010000000000", "ブドウ": "001000000000" ... etc]

    for token, vector := range onehotvectors {
        fmt.Println(token, vector)
    }

```
