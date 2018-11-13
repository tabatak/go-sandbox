package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	defer func() {
		fmt.Println("停止前に実行したい処理を実行")
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	done := make(chan interface{})
	go doSomething(done)

	select {
	case <-done:
		fmt.Println("すべての処理が完了")
	case s := <-c:
		fmt.Printf("シグナルを受信: %v\n", s)
	}

}

func doSomething(done chan interface{}) {
	fmt.Println("メインの処理を開始")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d 件目の処理中...\n", i)
		time.Sleep(2 * time.Second)
	}
	done <- struct{}{}
}
