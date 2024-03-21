package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Будем использовать метод Рунге-Кутты
// Введем дополнительные переменные u1 = x, u2 = x', v1 = y, v2 = y'

// func rungeKutta(u1, u2, v1, v2, t, h float64) (float64, float64, float64, float64) {

// 	k1 := mat.NewVecDense(4, []float64{u2, v2, 0, 0})
// 	k2 := mat.NewVecDense(4, []float64{-5*u1 + 2*v1, 0, 0, 0})
// 	k3 := mat.NewVecDense(4, []float64{v2, 0, u2, 0})
// 	k4 := mat.NewVecDense(4, []float64{2*u1 - 8*v1, 0, 0, u2})

// 	deltaU1 := (k1.At(0, 0) + 2*k2.At(0, 0) + 2*k3.At(0, 0) + k4.At(0, 0)) * (h / 6)
// 	deltaU2 := (k1.At(1, 0) + 2*k2.At(1, 0) + 2*k3.At(1, 0) + k4.At(1, 0)) * (h / 6)
// 	deltaV1 := (k1.At(2, 0) + 2*k2.At(2, 0) + 2*k3.At(2, 0) + k4.At(2, 0)) * (h / 6)
// 	deltaV2 := (k1.At(3, 0) + 2*k2.At(3, 0) + 2*k3.At(3, 0) + k4.At(3, 0)) * (h / 6)

// 	u1 += deltaU1
// 	u2 += deltaU2
// 	v1 += deltaV1
// 	v2 += deltaV2

// 	return u1, u2, v1, v2
// }

func dxdt(x, y float64) float64 {
	return (-5 * x) + (2 * y)
}

func dydt(x, y float64) float64 {
	return (2 * x) - (8 * y)
}

func main() {
	// Создаем график
	p := plot.New()

	// Задаем параметры график
	p.Title.Text = "Задача 1.2"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Создаем срез точек для хранения значений x, y
	points := make(plotter.XYs, 0)

	// Задаем начальные условия
	x := 1.0
	y := 1.0
	h := 0.01     // Шаг
	tmax := 100.0 // Максимальное время

	// Вычисляем значения точек
	for t := 0.0; t < tmax; t += h {
		x = x + h*dxdt(x, y)
		y = y + h*dydt(x, y)

		points = append(points, plotter.XY{X: x, Y: y})
	}

	// Создаем график с точками
	if err := plotutil.AddLines(p, points); err != nil {
		panic(err)
	}

	// Сохраняем график в файл
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "result.png"); err != nil {
		panic(err)
	}
}
