package students

import "fmt"

type Student struct {
	Name string
	age  int
}

func (s Student) PrintData() {
	fmt.Println(s.Name, s.age)
}

func (s Student) ChangeAge(age_ int) {
	s.age = age_
	fmt.Println("Age is changed")
}


// type Task struct {
// 	summary string
// 	description string
// 	deadline time.Time
// 	priority int
// }
// func (t Task) IsOverdue() bool {
// 	now := time.Now()
// 	if now.After(t.deadline) {
// 		return true
// 	} else {
// 		return false
// 	}
// }
// func (t Task) IsTopPriority() bool {
// 	if t.priority < 4 {
// 		return false
// 	} else {
// 		return true
// 	}
// }


// type Note struct {
// 	title string
// 	text string
// }


// type ToDoList struct {
// 	name string
// 	tasks []Task
// 	notes []Note
// }
// func (tdl ToDoList) TasksCount() int {
// 	return len(tdl.tasks)
// }
// func (tdl ToDoList) NotesCount() int {
// 	return len(tdl.notes)
// }
// func (tdl ToDoList) CountTopPrioritiesTasks() int {
// 	c := 0
// 	for _, tsk := range tdl.tasks {
// 		if tsk.IsTopPriority() {
// 			c++
// 		}
// 	}
// 	return c
// }
// func (tdl ToDoList) CountOverdueTasks() int {
// 	c := 0
// 	for _, tsk := range tdl.tasks {
// 		if tsk.IsOverdue() {
// 			c++
// 		}
// 	}
// 	return c
// }

