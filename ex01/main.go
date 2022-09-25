package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Action struct {
	actions []string
	Human
}

func (h *Human) getName() string {
	return h.name
}

func (h *Human) getAge() int {
	return h.age
}

func (h *Human) isOver18() bool {
	return h.age >= 18
}

func (a *Action) printActions() {
	fmt.Println(a.actions)
}

func main() {
	//инициализируем структуры human и action
	acts := []string{"Walk", "Read", "Seat"}
	human := Human{
		name: "Testovich",
		age:  19,
	}
	action := Action{
		actions: acts,
		Human:   human,
	}

	//вызываем методы Human
	fmt.Println(action.isOver18())
	fmt.Println(action.getName())
	fmt.Println(action.getAge())
	//вызываем методы Action
	action.printActions()
}
