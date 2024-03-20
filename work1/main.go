package main

import (
	"fmt"
	"os"

	"github.com/wcharczuk/go-chart/v2"
)

func main() {

	C := 10.0  // Емкость
	R := 100.0 // Споротивление
	e := 100.0 // ЭДС источника
	U := 0.0   // Начальное значение
	dt := 0.1  // Шаг

	var time []float64   // Массив со временем результатов
	var result []float64 // Массив с результатами

	// С начального времени до 10, шаг 0.1
	for t := 0.0; t <= 10.0; t += dt {
		k1 := dt * ((e - U) / (R * C))
		k2 := dt * ((e - (U + k1/2)) / (R * C))
		k3 := dt * ((e - (U + k2/2)) / (R * C))
		k4 := dt * ((e - (U + k3)) / (R * C))

		U += (k1 + 2*k2 + 2*k3 + k4) / 6 // Решение методом Рунге-Кутты

		time = append(time, t)
		result = append(result, U)
	}

	for i := 0; i < len(time); i++ { // Вывод данных на экран
		fmt.Printf("t = %f, U = %f\n", time[i], result[i])
	}

	// Строим наш график
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: time,
				YValues: result,
				Style: chart.Style{
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
			},
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top:  20,
				Left: 20,
			},
		},
		XAxis: chart.XAxis{
			Name:      "Время",
			Style:     chart.Shown(),
			NameStyle: chart.Shown(),
		},
		YAxis: chart.YAxis{
			Name:      "Напряжение",
			Style:     chart.Shown(),
			NameStyle: chart.Shown(),
		},
	}

	f, _ := os.Create("result.png")
	defer f.Close()
	if err := graph.Render(chart.PNG, f); err != nil {
		fmt.Print(err)
	}
}
