package processor

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type Result struct {
	Id    string
	Value int
}

type Input struct {
	Id   string
	Op   string
	Val1 int
	Val2 int
}

func parser(data []byte) (Input, error) {
	// parse the data
	lines := bytes.Split(data, []byte("\n"))
	// each entry is line 1 id, line 2 operator, line 3 num 1, line 4 num2
	id := string(lines[0])
	op := string(lines[1])
	val1, err := strconv.Atoi(string(lines[2]))
	if err != nil {
		return Input{}, err
	}
	val2, err := strconv.Atoi(string(lines[3]))
	if err != nil {
		return Input{}, err
	}
	return Input{
		Id:   id,
		Op:   op,
		Val1: val1,
		Val2: val2,
	}, nil
}

func DataProcessor(in <-chan []byte, out chan<- Result) {
	for data := range in {
		input, err := parser(data)
		if err != nil {
			continue
		}
		var calc int
		switch input.Op {
		case "+":
			calc = input.Val1 + input.Val2
		case "-":
			calc = input.Val1 - input.Val2
		case "*":
			calc = input.Val1 * input.Val2
		case "/":
			// 0で割ろうとした場合は何もしない
			if input.Val2 == 0 {
				continue
			}
			calc = input.Val1 / input.Val2
		default:
			continue
		}
		// sum numbers in the data
		// write to another channel
		result := Result{
			Id:    input.Id,
			Value: calc,
		}
		out <- result
	}
	close(out)
}

func WriteData(in <-chan Result, w io.Writer) {
	for r := range in {
		// write the output data to writer
		// each line is id:result
		w.Write([]byte(fmt.Sprintf("%s:%d\n", r.Id, r.Value)))
	}
}
