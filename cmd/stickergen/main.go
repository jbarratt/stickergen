package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/jbarratt/stickergen/render"
	"github.com/urfave/cli"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.UintFlag{Name: "rows, r", Value: 50},
		cli.UintFlag{Name: "cols, c", Value: 72},
		cli.UintFlag{Name: "cellsize, s", Value: 35},
		cli.StringFlag{Name: "outfile, o", Value: "output.png"},
		cli.StringFlag{Name: "palette, p", Value: "random", Usage: "Either a palette name (sendgrid, random) or a pair of hex colors"},
	}

	app.Action = func(c *cli.Context) error {
		output, _ := os.Create(c.String("outfile"))

		palette := c.String("palette")
		c1, c2 := palette, ""
		if colors := strings.Split(palette, ","); len(colors) > 1 {
			c1 = colors[0]
			c2 = colors[1]
		}

		err := render.GenerateImage(c.Uint("rows"), c.Uint("cols"), c.Uint("cellsize"), c1, c2, output)
		output.Close()

		if err != nil {
			return cli.NewExitError(err, 1)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
