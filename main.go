package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/gopherjs/gopherjs/js"
	"github.com/jbarratt/stickergen/render"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	js.Global.Set("generateImage", func(rows int, cols int, size int, c1 string, c2 string) string {
		output := new(bytes.Buffer)
		_ = render.GenerateImage(uint(rows), uint(cols), uint(size), c1, c2, output)
		b64Png := base64.StdEncoding.EncodeToString(output.Bytes())
		return fmt.Sprintf("data:image/png;base64,%s\n", b64Png)
	})
}
