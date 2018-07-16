package render

import (
	"errors"
	"fmt"
	"image/color"
	"io"
	"math/rand"

	"github.com/fogleman/gg"
	colorful "github.com/lucasb-eyer/go-colorful"
)

func loadPalette(c1 string, c2 string) []color.Color {
	var color1 colorful.Color
	var color2 colorful.Color
	var err error

	if c1 == "sendgrid" {
		return []color.Color{
			color.RGBA{130, 227, 247, 255},
			color.RGBA{159, 225, 240, 255},
			color.RGBA{0, 129, 233, 255},
			color.RGBA{0, 158, 223, 255},
			color.RGBA{0, 181, 232, 255},
			color.RGBA{0, 191, 233, 255},
			color.RGBA{0, 164, 213, 255},
			color.RGBA{0, 160, 224, 230},
		}
	} else if c1 == "random" {
		color1 = colorful.WarmColor()
		color2 = colorful.HappyColor()
	} else {
		color1, err = colorful.Hex(c1)
		if err != nil {
			fmt.Println("First color not a hex string (e.g. #00AA00), replacing with random warm color")
			color1 = colorful.WarmColor()
		}

		color2, err = colorful.Hex(c2)
		if err != nil {
			fmt.Println("Second color not a hex string (e.g. #00AA00), replacing with random happy color")
			color2 = colorful.HappyColor()
		}
	}

	// pal := colorful.FastHappyPalette(8)
	count := 7
	colors := make([]color.Color, count, count)
	for i := 0; i < count; i++ {
		colors[i] = color.Color(color1.BlendHcl(color2, float64(i)/float64(count-1)).Clamped())
	}
	return colors
}

func GenerateImage(rows uint, cols uint, size uint, c1 string, c2 string, w io.Writer) error {

	colors := loadPalette(c1, c2)
	colorCnt := len(colors)

	grid := make([][]color.Color, rows)
	for i := uint(0); i < rows; i++ {
		grid[i] = make([]color.Color, cols)
	}

	dc := gg.NewContext(int(size*cols), int(size*rows))
	dc.SetColor(color.RGBA{255, 255, 255, 255})
	dc.DrawRectangle(0, 0, float64(size*cols), float64(size*rows))
	dc.Fill()

	// iterate through every cell, row by row
	for row := uint(0); row < rows; row++ {
		for col := uint(0); col < cols; col++ {

			// Select a color for this cell; start with a random option
			offset := rand.Intn(colorCnt)
			var paintColor color.Color
			metConstraint := true

			// Starting
			for x := 0; x < colorCnt; x++ {
				idx := (x + offset) % colorCnt
				paintColor = colors[idx]
				// check the cells to the left of, and above, the new cell
				if col != 0 && paintColor == grid[row][col-1] {
					continue
				}
				if row != 0 && paintColor == grid[row-1][col] {
					continue
				}
				if x == (colorCnt - 1) {
					metConstraint = false
				}
				break
			}
			if !metConstraint {
				return errors.New("Unable to make placement rule, too few colors")
			}
			grid[row][col] = paintColor
			dc.SetColor(paintColor)
			dc.DrawRectangle(float64(col*size), float64(row*size), float64(size), float64(size))
			dc.Fill()
		}
	}
	return dc.EncodePNG(w)

}
