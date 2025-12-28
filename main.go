package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/crazy3lf/colorconv"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	width  int
	height int
	image  *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//ebitenutil.DebugPrint(screen, "Hello, World!")
	screen.DrawImage(g.image, nil)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}

func main() {
	if len(os.Args) == 2 {
		path := os.Args[1]

		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			wd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			_, err = os.Stat(wd + path)
			if err == nil {
				path = wd + path
			}
		}

		content_byte, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(err)
		}

		content := string(content_byte)

		index := strings.Index(content, ";")
		width_string := content[:index]
		width, err := strconv.Atoi(width_string)

		if err != nil {
			log.Fatal(err)
		}

		content = content[index+1:]

		index = strings.Index(content, ";")
		height_string := content[:index]
		height, err := strconv.Atoi(height_string)

		if err != nil {
			log.Fatal(err)
		}

		content = content[index+1:]

		var hex_slice []string

		for i := 0; i+6 <= len(content); i += 6 {
			hex_slice = append(hex_slice, content[i:i+6])
		}

		image := ebiten.NewImage(width, height)

		i := 0
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if i < len(hex_slice) {
					c, err := colorconv.HexToColor("#" + hex_slice[i])
					if err != nil {
						log.Fatal(err)
					}
					image.Set(x, y, c)
					i++
				}
			}
		}

		ebiten_image := ebiten.NewImageFromImage(image)

		ebiten.SetWindowSize(width, height)
		ebiten.SetWindowTitle("Hello, World!")
		if err := ebiten.RunGame(&Game{
			width:  width,
			height: height,
			image:  ebiten_image,
		}); err != nil {
			log.Fatal(err)
		}
	}
}
