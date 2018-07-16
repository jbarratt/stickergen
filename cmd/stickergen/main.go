package main

import (
	"log"
	"math/rand"
	"os"
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
		cli.StringFlag{Name: "palette, p", Value: "random"},
	}

	app.Action = func(c *cli.Context) error {
		output, _ := os.Create(c.String("outfile"))
		err := render.GenerateImage(c.Uint("rows"), c.Uint("cols"), c.Uint("cellsize"), c.String("palette"), "", output)
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
