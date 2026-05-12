package main

import (
	"errors"
	"fmt"
	"strconv"
)

func add(i int, j int) (int, error) { return i + j, nil } //liststart1

func sub(i int, j int) (int, error) { return i - j, nil }

func mul(i int, j int) (int, error) { return i * j, nil }

func div(i int, j int) (int, error) { //listend1
	if j == 0 {
		return 0, errors.New("0で割ることはできません")
	}
	return i / j, nil
}

var opMap = map[string]func(int, int) (int, error){ // 「文字列→関数」のマップ  //liststart2
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
} //listend2

func main() {
	expressions := [][]string{ // 例題（計算する式）  //liststart3
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "/", "0"},
		[]string{"2", "%", "3"},
		[]string{"two", "+", "three"},
		[]string{"2", "+", "three"},
		[]string{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 { // 演算子と被演算子の合計個数のチェック
			fmt.Print(expression, " -- 不正な式です\n")
			continue
		}
		p1, err := strconv.Atoi(expression[0]) // 1番目の被演算子（oPerand）のチェック
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		op := expression[1] // 演算子（OPerator）のチェック
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Print(expression, " -- ", "定義されていない演算子です: ", op, "\n")
			continue
		}
		p2, err := strconv.Atoi(expression[2]) // 2番目の被演算子のチェック
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		result, err := opFunc(p1, p2) // 実際の計算
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		fmt.Print(expression, " → ", result, "\n")
	} //listend3
}
