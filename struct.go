package main

import "fmt"

// タスクの定義
type Task struct {
	ID     int
	Detail string
	done   bool // 小文字は閉じたスコープ
}

// String() という振る舞いを規定するインターフェース
type Stringer interface {
	String() string
}

// コンストラクタを定義出来ないので自前でメソッドを用意
func newTask(id int, detail string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
		done:   false,
	}
	return task
}

// レシーバをポインタで受け取る例
func (task *Task) Finish() {
	task.done = true
}

// レシーバを値で受け取る例
// stringに変換するメソッド
func (task Task) String() string {
	str := fmt.Sprintf("[%t] %d:%s", task.done, task.ID, task.Detail)
	return str
}

func Print(stringer Stringer) {
	fmt.Println(stringer.String())
}

// どのような型でも受け付けるinterfaceの実装
func Do(e interface{}) {
	// do something
}

func main() {
	task := newTask(1, "Goの練習")
	task.Finish()
	Print(task)
}
