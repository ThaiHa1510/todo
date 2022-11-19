package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ThaiHa1510/todo"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "Thêm công việc")
	complete := flag.Int("complete", 0, "Đánh dấu công việc hoàn thành")
	del := flag.Int("delete", 0, "Xóa công việc")
	list := flag.Bool("list", false, "Hiển thị")
	flag.Parse()
	todos := &todo.Todo{}
	defer save(todos)
	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		todos.Add(task)
		todos.Print()
	case *complete > 0:
		err := todos.Complete(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		todos.Print()
	case *del > 0:
		err := todos.Delete(*del)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		todos.Print()
	case *list:
		todos.Print()
	default:
		todos.Print()
	}

}

func save(todos *todo.Todo) {
	todos.Store(todoFile)
}
func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	scaner := bufio.NewScanner(r)
	scaner.Scan()
	if err := scaner.Err(); err != nil {
		return "", err
	}
	text := scaner.Text()

	if len(text) == 0 {
		return "", errors.New("Vui lòng nhập tên task")
	}
	return text, nil

}
