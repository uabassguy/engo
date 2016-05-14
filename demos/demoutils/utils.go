package demoutils

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"engo.io/ecs"
	"engo.io/engo"
)

type Background struct {
	ecs.BasicEntity
	engo.RenderComponent
	engo.SpaceComponent
}

// NewBackground creates a background of colored tiles - might not be the most efficient way to do this
// It gets added to the world as well, so we won't return anything.
func NewBackground(world *ecs.World, width, height int, colorA, colorB color.Color) *Background {
	rect := image.Rect(0, 0, width, height)

	img := image.NewNRGBA(rect)
	for i := rect.Min.X; i < rect.Max.X; i++ {
		for j := rect.Min.Y; j < rect.Max.Y; j++ {
			if i%40 > 20 {
				if j%40 > 20 {
					img.Set(i, j, colorA)
				} else {
					img.Set(i, j, colorB)
				}
			} else {
				if j%40 > 20 {
					img.Set(i, j, colorB)
				} else {
					img.Set(i, j, colorA)
				}
			}
		}
	}

	bgTexture := engo.NewImageObject(img)

	bg := &Background{BasicEntity: ecs.NewBasic()}
	bg.RenderComponent = engo.RenderComponent{Drawable: engo.NewTexture(bgTexture)}
	bg.SpaceComponent = engo.SpaceComponent{
		Position: engo.Point{0, 0},
		Width:    float32(width),
		Height:   float32(height),
	}

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *engo.RenderSystem:
			sys.Add(&bg.BasicEntity, &bg.RenderComponent, &bg.SpaceComponent)
		}
	}

	return bg
}

// Creates a background image from a png file
// Probably should be updated to return a slice that is only as large 
// as the window so it isn't rendering the whole image on each frame
func NewBackgroundImage(world *ecs.World, src string) *Background {

	file, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	background, err := png.Decode(file)
	if err != nil {
		log.Fatal(os.Stderr, "%s: %v\n", src, err)
	}

	b := background.Bounds()

	log.Println(b)

	rect := image.Rect(0, 0, b.Max.X, b.Max.Y)
	img := image.NewNRGBA(rect)

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			oldPixel := background.At(x, y)
			r, g, b, a := oldPixel.RGBA()
			pixel := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			img.Set(x, y, pixel)
		}
	}

	bgTexture := engo.NewImageObject(img)

	bg := &Background{BasicEntity: ecs.NewBasic()}
	bg.RenderComponent = engo.RenderComponent{Drawable: engo.NewTexture(bgTexture)}
	bg.SpaceComponent = engo.SpaceComponent{
		Position: engo.Point{0, 0},
		Width:    float32(b.Max.X),
		Height:   float32(b.Max.Y),
	}

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *engo.RenderSystem:
			sys.Add(&bg.BasicEntity, &bg.RenderComponent, &bg.SpaceComponent)
		}
	}

	return bg
}
