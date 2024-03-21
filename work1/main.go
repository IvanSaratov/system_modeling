package main

import (
	"fmt"
	"os"

	"github.com/wcharczuk/go-chart/v2"
)

func dUdt(e, U, R, C float64) float64 {
	return (e - U) / (R * C)
}

func main() {

	C := 10.0       // Емкость
	R := 100.0      // Споротивление
	e := 100.0      // ЭДС источника
	U := 0.0        // Начальное значение
	dt := 100.0     // Шаг
	tmax := 10000.0 // Максимальное время

	var time []float64   // Массив со временем результатов
	var result []float64 // Массив с результатами

	for t := dt; t <= tmax; t += dt {
		tmp := U + dt*dUdt(e, U, R, C)
		time = append(time, t)
		result = append(result, tmp)

		U = tmp
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
