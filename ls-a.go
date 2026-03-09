package main

import (
	"fmt"
	"os"
	"flag"
	"strings"
)

func main() {
	dir := "."

	// flagはコマンドラインオプションを解析するための標準ライブラリ
	// -a, -m, -Xなどがflag
	// flag.Bool(name, default, description)
	// -aというフラグを定義, -aがついたらtrue
	showAll := flag.Bool("a", false, "隠しファイルも表示する") // *boolつまりboolポインタを返す
	flag.Parse() // これがあることでフラグが読み込まれる

	fmt.Println("flag.Args()", flag.Args()) // flag.Args()はフラグ以外の引数を返す

	args := flag.Args()
	if len(args) > 0 {
		dir = args[0]
		fmt.Println("args[0]:", args[0])
		fmt.Println("")
	}
	
	entries, err := os.ReadDir(dir) // osのReadDir隠しファイルも表示
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, entry := range entries {
		// strings.HasPrefixは文字列が指定の文字で始まるか
		if !*showAll && strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		fmt.Println(entry.Name())
	}
}
