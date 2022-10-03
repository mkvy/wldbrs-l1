package main

import "fmt"

type Set interface {
	AddToSet(val string) bool
	PrintSet()
}

// множество с поиском элемента при добавлении
type DefaultSet struct {
	vals []string
}

func CreateDefaultSet(strs []string) *DefaultSet {
	s := new(DefaultSet)
	for _, v := range strs {
		s.AddToSet(v)
	}
	return s
}

func (s *DefaultSet) AddToSet(val string) bool {
	//возвращает true, если элемента не было в множестве. если уже был, возвращает false
	for _, v := range s.vals {
		if v == val {
			return false
		}
	}
	s.vals = append(s.vals, val)
	return true
}

func (s *DefaultSet) PrintSet() {
	fmt.Print("{ ")
	for _, v := range s.vals {
		fmt.Printf("%s, ", v)
	}
	fmt.Print("}\n")
}

// множество на основе хеш-таблицы
type HashSet struct {
	vals map[string]bool
}

func CreateHashSet(strs []string) *HashSet {
	s := new(HashSet)
	s.vals = make(map[string]bool)
	for _, v := range strs {
		s.AddToSet(v)
	}
	return s
}

func (s *HashSet) AddToSet(val string) bool {
	//возвращает true, если элемента не было в множестве. если уже был, возвращает false
	if _, ok := s.vals[val]; !ok {
		s.vals[val] = true
		return true
	}
	return false
}

func (s *HashSet) PrintSet() {
	fmt.Print("{ ")
	for v, _ := range s.vals {
		fmt.Printf("%s, ", v)
	}
	fmt.Print("}\n")
}

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("Array: ", strs)
	var set Set = CreateDefaultSet(strs)
	fmt.Print("Def Set: ")
	set.PrintSet()
	set = CreateHashSet(strs)
	fmt.Print("Hash Set: ")
	set.PrintSet()
}
