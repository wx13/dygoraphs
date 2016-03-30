package main

import (
	"fmt"
	"github.com/wx13/dygoraphs"
)

func LinePlot() {

	fmt.Println("<h2>Line Plot</h2>")

	g := dygoraphs.NewGraph()
	g.Style = "width:500px; height:300px;"
	cols := []string{"time", "height", "weight"}
	data := [][]float64{
		[]float64{0, 1, 2, 3, 4, 5},
		[]float64{100, 102, 108, 109, 110, 110},
		[]float64{151, 149, 147, 150, 155, 160},
	}
	g.Add(cols, data)
	fmt.Println(g.Plot())

}

func ErrorBars() {

	fmt.Println("<h2>Error Bars</h2>")

	g := dygoraphs.NewGraph()
	g.Style = "width:500px; height:300px;"
	cols := []string{"time", "height", "weight"}
	data := [][]float64{
		[]float64{0, 1, 2, 3, 4, 5},
		[]float64{100, 102, 108, 109, 110, 110},
		[]float64{10, 10, 8, 9, 3, 10},
		[]float64{151, 149, 147, 150, 155, 160},
		[]float64{20, 25, 15, 10, 25, 20},
	}
	g.Options = "{errorBars: true}"
	g.Add(cols, data)
	fmt.Println(g.Plot())

}

func main() {

	fmt.Println("<html><head><title></title>")
	fmt.Println(dygoraphs.Script)
	fmt.Println("</head><body>")

	LinePlot()
	ErrorBars()

	fmt.Println("</body></html>")

}
