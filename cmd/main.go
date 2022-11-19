package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add new todo")
	complete := flag.Int("complete",0,"mark a todo as completed")
	del  := flag.Int("delete",0,"delete a todo")
	list := flag.Bool("list",false,"list all todos")

	flag.Parse()
	todos := &todo.Todo;

	if err:= todos.Load(todoFile),err != nil{
		fmt.Fprintln(os.Stderr,err.Error())
	}
	fmt.Println(*add)
}
