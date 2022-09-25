package main

import "fmt"

// через использование хеш-таблиц
func hashSetIntersect(s1 []string, s2 []string) (s3 []string) {
	hash := map[string]bool{}
	//добавляем в хеш таблицу наличие элемента из первого множества
	for _, v := range s1 {
		hash[v] = true
	}
	for _, v := range s2 {
		//проходим по второму; в случае, если оно содержится в первом, добавляем в выходное множество;
		//чтобы не было повторений ставим наличие false в мапе
		if hash[v] {
			hash[v] = false
			s3 = append(s3, v)
		}
	}
	return s3
}

func searchIntersect(s1 []string, s2 []string) (s3 []string) {
	//итерируемся по более длинному, ищем в более коротком
	if len(s1) > len(s2) {
		for _, v := range s1 {
			//проверяем есть ли элемент во втором массиве и проверяем на дубликат в выходном множестве
			if findVal(s2, v) && !findVal(s3, v) {
				s3 = append(s3, v)
			}
		}
	} else {
		for _, v := range s2 {
			if findVal(s1, v) && !findVal(s3, v) {
				s3 = append(s3, v)
			}
		}
	}

	return s3
}

// вспомогательная функция для поиска значения в массиве
func findVal(s []string, val string) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}
	return false
}

func main() {
	//первый способ: через использование хеш-таблиц
	s1 := []string{"pink", "brown", "red", "blue", "red", "blue"}
	s2 := []string{"brown", "red", "yellow", "pink", "red"}
	s3 := hashSetIntersect(s1, s2)
	fmt.Println(s3)

	//второй способ: через перебор и поиск
	s3 = searchIntersect(s1, s2)
	fmt.Println(s3)

}
