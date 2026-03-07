package main
/* 
packageによってコードを機能ごとにグループ化することができる
ex. math, user, authなど

user.Get(), auth.Get()のように区別できる

importで再利用する

package mainはGoのプログラムのエントリーポイント（実行アプリ）になる
特別なパッケージ
*/

import "fmt" // fmt パッケージを読み込んでいるのでfmt.Println()とかける

func main() { // プログラム開始時に呼ばれる入り口の関数
	fmt.Println("hello, world")
}
