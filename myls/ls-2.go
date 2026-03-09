package main

import (
	"fmt"
	"os"
)

func main() {
	dir := "."
	// os.Argsはコマンドライン引数の配列
	// go run main.go test の時[0]がmain.go, [1]がtest
	// go run するとどこかにバイナリが生成され、同じコードならキャッシュを参照する

	// なぜ > 1なのかは実際に出力すればわかる
	if len(os.Args) > 1 { // os.Argsにはまずファイル名が入っているので1より大きい場合にdir に代入する
		dir = os.Args[1]
	}
	
	for i, arg := range os.Args {
		fmt.Printf("os.Args[%v]: %v\n", i, arg)
	}
	fmt.Println("") // print line

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}
