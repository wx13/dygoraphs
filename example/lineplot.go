package main

import (
	"fmt"
	"github.com/wx13/dygoraphs"
)

func main() {
	fmt.Println("<html><head><title></title>")
	fmt.Println(dygoraphs.Script)
	fmt.Println("</head><body>")
	fmt.Println("<h1>Line Plot</h1>")

	g := dygoraphs.NewGraph()
	cols := []string{"time", "height", "weight"}
	data := [][]float64{
		[]float64{0, 1, 2, 3, 4, 5},
		[]float64{100, 102, 108, 109, 110, 110},
		[]float64{151, 149, 147, 150, 155, 160},
	}
	g.Add(cols, data)
	g.Style = "width:600px; height:400px;"
	fmt.Println(g.Plot())

	fmt.Println("</body></html>")
}

