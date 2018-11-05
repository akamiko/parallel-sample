package main

import (
	"fmt"
	"testing"
	"time"
)

// runメソッドの時間を計測
func main() {
	result := testing.Benchmark(func(b *testing.B) { run() })
	fmt.Println(result.T)
}

// 2秒かかるプロセス３つを並列で実行
func run() {
	// それぞれのプロセス用にchannelを作成(boolean型でtrueなら処理終了)
	isFin1 := make(chan bool)
	isFin2 := make(chan bool)
	isFin3 := make(chan bool)

	fmt.Println("Start!!")
	// goroutineを生成して、サブスレッドで処理する。終了後、chに対してtrueを投げる
	go func() {
		process("A")
		isFin1 <- true
	}()
	go func() {
		process("B")
		isFin2 <- true
	}()
	go func() {
		process("C")
		isFin3 <- true
	}()

	// 全てのプロセスが終了するまでブロック
	<-isFin1
	<-isFin2
	<-isFin3

	// それぞれのchannelクローズ
	close(isFin1)
	close(isFin2)
	close(isFin3)

	fmt.Println("Finish!!")
}

// 2秒かかるプロセスを実行
func process(name string) {
	time.Sleep(2 * time.Second)
	fmt.Println(name)
}
