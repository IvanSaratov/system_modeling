package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Параметры нормального распределения
	mu1 := 10.0   // Среднее значение для первого элемента
	sigma1 := 2.0 // Стандартное отклонение для первого элемента
	mu2 := 12.0   // Среднее значение для второго элемента
	sigma2 := 3.0 // Стандартное отклонение для второго элемента
	mu3 := 15.0   // Среднее значение для третьего элемента
	sigma3 := 4.0 // Стандартное отклонение для третьего элемента

	numTrials := 10000
	totalDuration := 0.0
	totalSqrDuration := 0.0

	for i := 0; i < numTrials; i++ {
		// Генерируем длительность работы каждого элемента
		duration1 := rand.NormFloat64()*sigma1 + mu1
		duration2 := rand.NormFloat64()*sigma2 + mu2
		duration3 := rand.NormFloat64()*sigma3 + mu3

		// Находим минимальную из длительностей (т.е. первый отказавший элемент)
		minDuration := duration1
		if duration2 < minDuration {
			minDuration = duration2
		}
		if duration3 < minDuration {
			minDuration = duration3
		}

		// Обновляем общую длительность и квадрат длительности
		totalDuration += minDuration
		totalSqrDuration += minDuration * minDuration
	}

	// Находим среднюю длительность и среднее квадратическое отклонение
	meanDuration := totalDuration / float64(numTrials)
	meanSqrDuration := totalSqrDuration / float64(numTrials)
	stdDev := meanSqrDuration - meanDuration*meanDuration

	fmt.Println("Математическое ожидание длительности безотказной работы системы:", meanDuration)
	fmt.Println("Среднее квадратическое отклонение длительности безотказной работы системы:", stdDev)
}
