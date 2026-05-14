package main

type Team struct {
	Name        string
	PlayerNames []string
}

type League struct {
	Teams []Team
	Wins  map[string]int // チーム名をキー、勝利数を値とするマップ
}
