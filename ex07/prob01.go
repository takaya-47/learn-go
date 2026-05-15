package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
)

type Ranker interface {
	Ranking() []string
}

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team // チーム名をキー、Team構造体を値とするマップ
	Wins  map[string]int  // チーム名をキー、勝利数を値とするマップ
}

func (l *League) MatchResult(team1 string, score1 int, team2 string, score2 int) {
	if _, ok := l.Teams[team1]; !ok {
		log.Printf("チーム %s はリーグに存在しません\n", team1)
		return
	}

	if _, ok := l.Teams[team2]; !ok {
		log.Printf("チーム %s はリーグに存在しません\n", team2)
		return
	}

	if score1 > score2 {
		l.Wins[team1]++
		return
	}

	if score1 == score2 {
		log.Printf("引き分け: %s vs %s\n", team1, team2)
		return
	}

	if score1 < score2 {
		l.Wins[team2]++
	}
}

func (l *League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))
	for k := range l.Teams {
		names = append(names, k)
	}

	slices.SortFunc(names, func(a, b string) int {
		if l.Wins[a] > l.Wins[b] {
			return -1
		}
		if l.Wins[a] < l.Wins[b] {
			return 1
		}
		return 0
	})

	return names
}

func RankPrinter(r Ranker, w io.Writer) {
	ranking := r.Ranking()
	for _, team := range ranking {
		fmt.Fprintln(w, team)

		// 以下でも同じ結果。
		// io.WriteString(w, team)
		// io.WriteString(w, "\n")
	}
}

func main() {
	l := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Italy": {
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"France": {
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"India": {
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Nigeria": {
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	// France : 2勝
	// Nigeria : 2勝
	// Italy : 1勝
	// India : 1勝
	l.MatchResult("Italy", 50, "France", 70)
	l.MatchResult("India", 85, "Nigeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Nigeria", 110)
	l.MatchResult("Italy", 65, "Nigeria", 70)
	l.MatchResult("France", 95, "India", 80)
	results := l.Ranking()
	fmt.Println(results)

	RankPrinter(&l, os.Stdout)
}
