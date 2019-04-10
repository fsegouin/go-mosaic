package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/ajstarks/svgo"
)

func main() {
	http.Handle("/", http.HandlerFunc(mosaic))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func mosaic(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	rand.Seed(time.Now().UTC().UnixNano())
	width := 3000
	height := 2000
	squareWidth := 50
	canvas := svg.New(w)
	canvas.Start(width, height)
	canvas.Gtransform("rotate(-45)")
	canvas.Gtransform("translate(-800, 0)")

	// Colors

	colors := [][]int{}

	color1 := []int{65, 164, 179}
	color2 := []int{253, 253, 253}
	color3 := []int{133, 44, 110}
	color4 := []int{251, 44, 92}
	color5 := []int{252, 224, 1}

	colors = append(colors, color1, color2, color3, color4, color5)

	//r := rand.Intn(20)
	r := 60
	s := 60
	for i := 0; i < r; i++ {
		for j := 0; j < s; j++ {
			color := rand.Intn(5)
			fill := canvas.RGBA(colors[color][0],
				colors[color][1],
				colors[color][2],
				1)
			canvas.Square((squareWidth-1)*i, (squareWidth-1)*j, squareWidth, fill)
		}
	}
	canvas.End()
}
