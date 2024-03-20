package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/wcharczuk/go-chart/v2"
)

func main() {
	// Параметры для первого способа
	probRepair1 := 0.2
	probPack1 := 0.8

	// Параметры для второго способа
	probRepair2 := 0.2
	probPack2 := 0.8

	// Количество телевизоров
	numTVs := []int{20, 50, 500, 1000, 10000}

	result := []chart.Value{}

	// Моделирование и оценка вероятности упаковки для обоих способов
	for _, n := range numTVs {
		pack1 := simulateTVs(n, probRepair1, probPack1)
		pack2 := simulateTVs(n, probRepair2, probPack2)
		// Выводим вероятности упаковки
		fmt.Printf("Для %d телевизоров: способ 1 - %.5f, способ 2 - %.5f\n", n, pack1, pack2)
		result = append(result, chart.Value{Value: pack1, Label: strconv.Itoa(n)})
		result = append(result, chart.Value{Value: pack2, Label: strconv.Itoa(n)})
	}

	// Строим наш график
	graph := chart.BarChart{
		Title: "Вероятность упаковки телевизора",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		Bars:     result,
	}

	f, _ := os.Create("result.png")
	defer f.Close()
	if err := graph.Render(chart.PNG, f); err != nil {
		panic(err)
	}
}

// Моделирование процесса проверки телевизоров
func simulateTVs(numTVs int, probRepair, probPack float64) float64 {
	packedTVs := 0
	for i := 0; i < numTVs; i++ {
		if rand.Float64() > probRepair { // Успешная проверка
			if rand.Float64() < probPack { // Упаковать
				packedTVs++
			}
		}
	}
	return float64(packedTVs) / float64(numTVs) // Возвращаем вероятность упаковки
}
