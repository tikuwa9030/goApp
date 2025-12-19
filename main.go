package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"io"
	"errors"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// go run・・・「ソースコードをその場でコンパイルして即実行」
// go build・・・「ソースコードをコンパイルしてバイナリを作るだけ」
// go get・・・「外部ライブラリをインストールするとき」

func main() {

	//配列は[]string{}の形式 []に長さ、{}に値
	pg := []string{"税込計算プログラム", "複利計算プログラム", "CSV読み書きプログラム", "ポインタ勉強用プログラム"}

	// for文は括弧でくくらない
	// iはカウンタ、vは配列の値を保持
	for i, v := range pg {
		fmt.Printf("%d: %s\n", i+1, v)
	}

	// Scanfの差異は値をポインタで参照
	fmt.Println("実行するアプリを選択してください。")
	selectNumber := 0

	// Scanfはあまり使わないほうがいいのが常識
	fmt.Scanln(&selectNumber)
	fmt.Println("選択された番号:", selectNumber)

	switch selectNumber {
	case 1:
		calculateTax()
	case 2:
		caluculateCompoundInterest()
	case 3:
		readCsv()
	case 4:
		pointerExample()
	default:
		fmt.Println("無効な番号です")
	}

}

func calculateTax() {
	var price int
	var priceWithTax float64
	const tax = 1.10

	// "_"はいらない値を捨てる変数（ゴミ箱）、nilは
	fmt.Println("価格を入力してください。")
	_, err := fmt.Scanln(&price)
	if err != nil {
		fmt.Println("数字を入力してください")
		return
	}

	priceWithTax = float64(price) * tax
	fmt.Printf("税込価格:%.0f\n", priceWithTax)

}

func caluculateCompoundInterest() {
	var principal int
	var rate float64
	var numberOfCal int
	var term int
	var result float64

	fmt.Println("元本:")
	fmt.Scanln(&principal)
	fmt.Println("年利率:")
	fmt.Scanln(&rate)
	fmt.Println("計算回数:")
	fmt.Scanln(&numberOfCal)
	fmt.Println("期間")
	fmt.Scanln(&term)

	result = float64(principal) * math.Pow(1+rate, float64(numberOfCal))

	fmt.Println("計算結果:", int(result))
}

func readCsv() {
	f, err := os.Open("ご利用明細_202512.csv")
	// ファイルが存在するか否か
	if err != nil {
		fmt.Println("開けない：", err)
		return
	}
	// deferは関数を抜ける直前に実行する
	defer f.Close()

	// Shift_JIS → UTF-8 に変換するReaderをかませる
	decodedReader := transform.NewReader(f, japanese.ShiftJIS.NewDecoder())

	r := csv.NewReader(decodedReader)
	r.FieldsPerRecord = -1 // 列数バラつき許容（クレカ明細系はバラつきがち）

	for {
		// recordに行ごとのデータを格納
		// Goのerrは失敗フラグではなく、状態通知（成功、終了、失敗）を伝えるもの
		record, err := r.Read()
		// 各行ごとのデータを読み込めるかどうか
		if err != nil {
			// ファイル内のデータがもう存在しないかどうか
			// err.Error() == "EOF" はやめた方がいい。エラー文字列比較は壊れやすい。Go では io.EOF を使うのが定番。
			// errors.Is(a, b)は"aはbというエラーか"という意味
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Println("csv読込エラー:", err)
			return
		}
		fmt.Println("行:", record)
	}
}

func pointerExample() {
	var i int
	// pはint型のポインタ変数、書き方は「*int」
	var p *int
	// pは変数iを参照するポインタ
	p = &i
	i = 5
	// *pはiを参照している
	fmt.Println(*p) // -> "5"
	// *pを通してiに10を代入
	*p = 10
	fmt.Println(i) // -> "10"
}
