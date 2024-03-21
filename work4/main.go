package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/fogleman/gg"
)

type Point struct {
	X float64
	Y float64
}

func isInsideTriangle(point Point, triangle [3]Point) bool {
	A := triangle[0]
	B := triangle[1]
	C := triangle[2]

	// Вычисляем вектора AB, BC и CA
	AB := Point{B.X - A.X, B.Y - A.Y}
	BC := Point{C.X - B.X, C.Y - B.Y}
	CA := Point{A.X - C.X, A.Y - C.Y}

	// Вычисляем вектора точки к вершинам треугольника
	AP := Point{point.X - A.X, point.Y - A.Y}
	BP := Point{point.X - B.X, point.Y - B.Y}
	CP := Point{point.X - C.X, point.Y - C.Y}

	// Вычисляем векторные произведения
	crossABP := AB.X*AP.Y - AB.Y*AP.X
	crossBCP := BC.X*BP.Y - BC.Y*BP.X
	crossCAP := CA.X*CP.Y - CA.Y*CP.X

	// Если точка находится с одной стороны от всех сторон треугольника,
	// то она внутри
	return (crossABP >= 0 && crossBCP >= 0 && crossCAP >= 0) || (crossABP < 0 && crossBCP < 0 && crossCAP < 0)
}

func randomPointInCircle(radius float64) Point {
	r := radius * math.Sqrt(rand.Float64())
	theta := rand.Float64() * 2 * math.Pi
	x := r * math.Cos(theta)
	y := r * math.Sin(theta)
	return Point{x, y}
}

func main() {
	// Количество случайных точек
	numPoints := 100
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Радиус круга
	cx, cy, radius := S/2.0, S/2.0, S/3.0

	dc.DrawCircle(cx, cy, radius)
	dc.SetRGB(1, 0, 0)
	dc.Stroke()

	x1, y1 := cx, cy-radius
	x2, y2 := cx+radius*math.Sin(math.Pi/3), cy+radius*math.Cos(math.Pi/3)
	x3, y3 := cx-radius*math.Sin(math.Pi/3), cy+radius*math.Cos(math.Pi/3)

	dc.DrawLine(x1, y1, x2, y2)
	dc.DrawLine(x2, y2, x3, y3)
	dc.DrawLine(x3, y3, x1, y1)
	dc.SetRGB(0, 255, 0)
	dc.Stroke()

	// Считаем количество точек, попавших внутрь треугольника
	countInside := 0
	for i := 0; i < numPoints; i++ {
		point := randomPointInCircle(radius)
		point.X += cx
		point.Y += cy

		if isInsideTriangle(point, [3]Point{
			{x1, y1}, {x2, y2}, {x3, y3},
		}) {
			countInside++
		}
		dc.DrawPoint(point.X, point.Y, 3.0)
	}

	dc.Stroke()
	dc.SavePNG("result.png")

	// Оцениваем вероятность
	probability := float64(countInside) / float64(numPoints)
	fmt.Printf("Вероятность: %f\n", probability)
}
