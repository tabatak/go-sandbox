package main

import (
	"errors"
	"log"
)

func main() {

	//switchの判定に使う値を取得
	// v, err := getValue()
	// if err != nil {
	// 	return
	// }

	// switch v {
	// case 1:
	// 	//1の場合の処理
	// 	log.Printf("proc of case 1")
	// case 2:
	// 	//1の場合の処理
	// 	log.Printf("proc of case 2")
	// default:
	// }

	// //switchの判定に使う値を取得
	// v, err := getValue()
	// if err != nil {
	// 	return
	// }

	//複数の戻り値を使うswitch
	switch v, err := getValue(); {
	case err != nil:
		log.Printf("somothing wrong. %v", err)
	case v == 1:
		//1の場合の処理
		log.Printf("proc of case 1")
	case v == 2:
		//1の場合の処理
		log.Printf("proc of case 2")
	default:
	}

}

func getValue() (int, error) {
	return 1, errors.New("error")
}
