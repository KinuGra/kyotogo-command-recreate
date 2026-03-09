package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 引数がなければerror
	dir := ""
	if (len(os.Args) < 2) {
		fmt.Fprintln(os.Stderr, "error: ファイルを指定してください")
		os.Exit(1)
	} else if (len(os.Args) == 3) {
		dir = os.Args[2]
	}
	
	data, err := os.ReadFile(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}	
