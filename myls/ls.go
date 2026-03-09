package main

import (
	"fmt"
	"os"
)

func main() {
	// カレントディレクトリのファイル一覧を取得
	entries, err := os.ReadDir(".") // Goはエラーを戻り値として返す

	fmt.Printf("entries: %v\n", entries)
	fmt.Printf("err: %v\n", err)
	fmt.Println("")
	
	// エラーチェック
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, entry := range entries {
		fmt.Println(entry.Name()) // ファイル名を表示
	}
}
