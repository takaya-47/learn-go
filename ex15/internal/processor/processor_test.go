package processor_test

import (
	"bytes"
	"strconv"
	"takaya-47/learn-go/ex15/internal/processor"
	"testing"
)

func TestDataProcessor(t *testing.T) {
	data := []struct {
		name     string
		operator string
		val1     int
		val2     int
		expected int
		wantErr  bool
	}{
		{"CALC_1", "+", 1, 2, 3, false},
		{"CALC_2", "-", 5, 3, 2, false},
		{"CALC_3", "*", 4, 6, 24, false},
		{"CALC_4", "/", 8, 2, 4, false},
		{"CALC_5", "&", 10, 15, 25, true}, // 演算子が不正
		{"CALC_6", "/", 8, 0, 0, true},    // 0で割る場合
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
			if d.wantErr {
				if ok {
					t.Error("Expected no result, but got one")
				}
				return
			}

			if !d.wantErr {
				if !ok {
					t.Fatal("Expected result, but got none")
				}

				if result.Id != d.name {
					t.Errorf("Expected name is %s, but got %s", d.name, result.Id)
				}
				if result.Value != d.expected {
					t.Errorf("Expected value is %d, but got %d", d.expected, result.Value)
				}
			}
		})
	}
}

func FuzzDataProcessor(f *testing.F) {
	data := []struct {
		name     string
		operator string
		val1     int
		val2     int
		expected int
		wantErr  bool
	}{
		{"CALC_1", "+", 1, 2, 3, false},
		{"CALC_2", "-", 5, 3, 2, false},
		{"CALC_3", "*", 4, 6, 24, false},
		{"CALC_4", "/", 8, 2, 4, false},
	}
	for _, d := range data {
		f.Add(d.name, d.operator, d.val1, d.val2)
	}

	f.Fuzz(func(t *testing.T, name string, operator string, val1 int, val2 int) {
		byteText := []byte(name + "\n" + operator + "\n" + strconv.Itoa(val1) + "\n" + strconv.Itoa(val2))
		in := make(chan []byte, 1)
		out := make(chan processor.Result, 1)

		go processor.DataProcessor(in, out)
		in <- byteText
		close(in)

		result, ok := <-out
		if !ok {
			t.Log("0 division or invalid operator, no result returned")
			return
		}

		got := result.Value
		switch operator {
		case "+":
			if want := val1 + val2; got != want {
				t.Errorf("want %d, got %d", want, got)
			}
		case "-":
			if want := val1 - val2; got != want {
				t.Errorf("want %d, got %d", want, got)
			}
		case "*":
			if want := val1 * val2; got != want {
				t.Errorf("want %d, got %d", want, got)
			}
		case "/":
			if want := val1 / val2; got != want {
				t.Errorf("want %d, got %d", want, got)
			}
		}
	})
}

func TestWriteData(t *testing.T) {
	// Arrange
	result := processor.Result{
		Id:    "CALC_1",
		Value: 3,
	}
	in := make(chan processor.Result, 1)
	var buf bytes.Buffer

	// Act
	in <- result
	close(in)
	processor.WriteData(in, &buf)

	// Assert
	expected := result.Id + ":" + strconv.Itoa(result.Value) + "\n"
	if got := buf.String(); got != expected {
		t.Errorf("Expected output is %s, but got %s", expected, got)
	}
}
