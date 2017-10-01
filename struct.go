package main

import "fmt"

type Task struct {
	ID     int
	Detail string
	done   bool // 小文字は閉じたスコープ
}

func (task Task) String() string {
	str := fmt.Sprintf("%d:%s", task.ID, task.Detail)
	return str
}

func main() {
	var task Task = Task{
		ID:     1,
		Detail: "goの練習",
		done:   true,
	}
	fmt.Println(task)
}
