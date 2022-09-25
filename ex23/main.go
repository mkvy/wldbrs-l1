package main

import "fmt"

func removeFromSlice(sl []int, index int) []int {
	if index > 0 && index < len(sl) {
		return append(sl[:index], sl[index+1:]...)
	} else {
		fmt.Println("Out of index")
		return nil
	}
}

// быстрое удаление, порядок не сохраняется
func fastRemove(sl []int, index int) []int {
	//последний элемент меняем на заменяемый
	sl[index] = sl[len(sl)-1]
	sl[len(sl)-1] = 0 //устанавливаем zero-value
	//возвращаем укороченный на 1 слайс
	return sl[:len(sl)-1]
}

func main() {
	sl := []int{1, 2, 3, 4, 5}
	fmt.Println("Исходный слайс: ", sl)
	fmt.Println(removeFromSlice(sl, 2))
	sl = []int{1, 2, 3, 4, 5}
	fmt.Println(fastRemove(sl, 2))
}
