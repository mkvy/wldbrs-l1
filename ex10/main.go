package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	datamap := map[int][]float32{}
	//приводим к int float, получаем значение нужной группы
	for _, v := range arr {
		datamap[int(math.Trunc(float64(v))/10)*10] = append(datamap[int(math.Trunc(float64(v))/10)*10], v)
	}

	fmt.Println(datamap)
}
