package main

// typeで既存の型を拡張して独自の型を定義
type ID int
type Priority int

func ProcessTask(id ID, priority Priority) {
	// 何らかの処理
}

func main() {
	var id ID = 1
	var priority Priority = 10
	ProcessTask(id, priority)
}
