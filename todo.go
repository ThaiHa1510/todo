package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

type Item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todo []Item

func (todo *Todo) Add(task string) {
	t := Item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*todo = append(*todo, t)
}

func (todo *Todo) Complete(index int) error {
	ls := *todo
	index--
	if index < 0 && index < len(ls) {
		return errors.New("Không tìm thấy công việc")
	}
	ls[index].Done = true
	ls[index].CompletedAt = time.Now()
	return nil
}

func (todo *Todo) Delete(index int) error {
	ls := *todo
	if index > 0 && index < len(ls) {
		return errors.New("Không tìm thấy công việc")
	}
	*todo = append(ls[:index-1], ls[index:]...)
	return nil
}

func (todo *Todo) Update(index int, task string) error {
	ls := *todo
	if index > 0 && index < len(ls) {
		return errors.New("Không tìm thấy công việc")
	}
	ls[index].Task = task
	return nil
}

func (todo *Todo) Store(filename string) error {
	data, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func (todo *Todo) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return err
	}
	err = json.Unmarshal(data, todo)
	return err
}

func (todo *Todo) Print() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignCenter, Text: "Created At"},
			{Align: simpletable.AlignCenter, Text: "Completed At"},
		},
	}
	var cells [][]*simpletable.Cell
	for idx, item := range *todo {
		idx++
		task := blue(item.Task)
		done := blue("no")
		complete := ""
		if item.Done {
			done = green("yes")
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			complete = item.CompletedAt.Format(time.RFC822)
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: complete},
		})

	}

	table.Body = &simpletable.Body{Cells: cells}
	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("Bạn có %d công việc cần hoàn thành", todo.CountPending()))},
		},
	}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}

func (todo *Todo) CountPending() int {
	var result = 0
	for _, item := range *todo {
		if !item.Done {
			result++
		}
	}
	return result
}
