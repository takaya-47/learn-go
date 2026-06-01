package main

import (
	"fmt"
	"math"
	"sync"
)

var squareRootMap func() map[int]float64 = sync.OnceValue(buildSquareRootMap)

func buildSquareRootMap() map[int]float64 {
	m := make(map[int]float64, 100_000)
	for i := 0; i < 100_000; i++ {
		m[i] = math.Sqrt(float64(i))
	}
	return m
}

func main() {
	// パッケージ変数のsquareRootMap（関数）を呼び出すと、初回はbuildSquareRootMapが呼び出されてマップが構築される。
	m := squareRootMap()

	// 以下性能検証用コード
	// start := time.Now()
	// for i := 0; i < 50000; i++ {
	// 	// 2回目以降はキャッシュされたマップが返されるので、ループ回数を増やしてもパフォーマンスに影響はない。
	// 	squareRootMap()
	// }
	// fmt.Println(time.Since(start))
	// 性能検証用コード。ここまで。

	for i := 0; i < len(m); i += 1_000 {
		fmt.Printf("%d %v\n", i, m[i])
	}
}
