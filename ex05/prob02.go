// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// )

// func main() {
// 	// go run ./prob02.go ./sample.txtのように実行するので、引数は以下のようになる。
// 	// os.Args[0] = "./prob02.go" => 実行対象のファイル名
// 	// os.Args[1] = "./sample.txt" => バイト数を知りたいファイル名
// 	if len(os.Args) < 2 {
// 		log.Fatal("引数が不正です")
// 		return
// 	}

// 	len, err := fileLen(os.Args[1])
// 	if err != nil {
// 		log.Fatalf("ファイル読み取りでエラー発生: %v", err)
// 	}
// 	fmt.Printf("File length: %v bytes\n", len)
// }

// func fileLen(fileName string) (int, error) {
// 	f, err := os.Open(fileName)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer f.Close()

// 	data := make([]byte, 2048) // 0が2048個入ったスライスを作成(長さが2048)
// 	total := 0
// 	for {
// 		count, err := f.Read(data) // ファイルから2048バイトずつ読み込む
// 		total += count
// 		if err != nil {
// 			if err == io.EOF {
// 				// ファイルの終わりに達した場合は、ループを抜ける
// 				break
// 			}
// 			return 0, err // 読み取り中にエラー発生時はバイト数を0にしてエラーを返す
// 		}
// 	}
// 	return total, nil
// }
