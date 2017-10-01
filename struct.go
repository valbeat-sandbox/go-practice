package main

import "fmt"

type Task struct {
	ID     int
	Detail string
	done   bool // 小文字は閉じたスコープ
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

func main() {
	task := newTask(1, "Goの練習")
	task.Finish()
	fmt.Println(task)
}
