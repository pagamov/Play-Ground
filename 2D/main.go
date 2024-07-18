package main

import (
	"image/color"
	"log"

	// "vector"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	Width  = 800
	Height = 600
)

type Game struct {
	lineArr []Line
}

type Line struct {
	x1, y1, x2, y2 float32
	color          color.Color
}

// type Ball struct {
// 	x, y, xVel, yVel float32
// 	color            color.Color
// }

type Vector struct {
	x, y float32
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < len(g.lineArr); i++ {
		vector.StrokeCircle(screen, g.lineArr[i].x1, g.lineArr[i].y1, 4, 2, color.Opaque, false)
		vector.StrokeLine(screen, g.lineArr[i].x1, g.lineArr[i].y1, g.lineArr[i].x2, g.lineArr[i].y2, 2, g.lineArr[i].color, false)
		vector.StrokeCircle(screen, g.lineArr[i].x2, g.lineArr[i].y2, 4, 2, color.Opaque, false)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}

func main() {
	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("Hello, World!")

	var center Vector = Vector{Width / 2, Height / 2}
	var length []float32 = []float32{70, 100, 120}

	var g Game = Game{[]Line{}}

	for i := 0; i < len(length); i++ {
		if i != 0 {
			g.lineArr = append(g.lineArr, Line{g.lineArr[i-1].x1, g.lineArr[i-1].y1, g.lineArr[i-1].x2, g.lineArr[i-1].y2 - length[i], color.Opaque})
		}
		g.lineArr = append(g.lineArr, Line{center.x, center.y, center.x, center.y - length[i], color.Opaque})
	}

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
