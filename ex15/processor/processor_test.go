package processor_test

// TODO: 0で割った時の異常系を追加し、プロダクトコードも修正

import (
	"strconv"
	"takaya-47/learn-go/ex15/processor"
	"testing"
)

func TestDataProcessor(t *testing.T) {
	data := []struct {
		name     string
		operator string
		val1     int
		val2     int
		expected int
	}{
		{"CALC_1", "+", 1, 2, 3},
		{"CALC_2", "-", 5, 3, 2},
		{"CALC_3", "*", 4, 6, 24},
		{"CALC_4", "/", 8, 2, 4},
		{"CALC_5", "&", 10, 15, 25}, // 演算子が不正
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			// Arrange
			byteText := []byte(d.name + "\n" + d.operator + "\n" + strconv.Itoa(d.val1) + "\n" + strconv.Itoa(d.val2))
			in := make(chan []byte, 1)
			out := make(chan processor.Result, 1)

			// Act
			// 普通に呼び出すとDataProcessorのforでブロックされてしまうため、inチャネルにデータを送信する部分に進めない。
			// よって、DataProcessorをゴルーチンで呼び出す
			go processor.DataProcessor(in, out)
			in <- byteText
			// inをクローズすることでDataProcessorのforループが終了し、関数自体も終了する（goroutineが終了するのでリーク対策になる。）
			close(in)

			// Assert
			result, ok := <-out
			// 正常系
			if ok {
				if result.Id != d.name {
					t.Errorf("Expected %s, got %s", d.name, result.Id)
				}
				if result.Value != d.expected {
					t.Errorf("Expected %d, got %d", d.expected, result.Value)
				}

				return
			}

			// 以下、異常系の検証
			if result.Id != "" && result.Value != 0 {
				t.Errorf("Expected empty result, but got %v, %v", result.Id, result.Value)
			}
		})
	}
}
