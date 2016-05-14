package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/demos/demoutils"
)

type DefaultScene struct{}

var (
	scrollSpeed float32 = 350

	// Size of the window
	cameraWidth  int = 400
	cameraHeight int = 400
)

func (*DefaultScene) Preload() {}

// Setup is called before the main loop is started
func (*DefaultScene) Setup(w *ecs.World) {
	w.AddSystem(&engo.RenderSystem{})

	// The most important line in this whole demo:
	w.AddSystem(engo.NewKeyboardScroller(scrollSpeed, engo.DefaultHorizontalAxis, engo.DefaultVerticalAxis))

	// Create a background image from a file (paths are relative)
	bg := demoutils.NewBackgroundImage(w, "assets/bg.png")

	// Create the boundaries of the world based on the size of the background image loaded
	engo.WorldBounds.Min = engo.Point{float32(cameraWidth), float32(cameraHeight)}
	engo.WorldBounds.Max = engo.Point{float32(bg.SpaceComponent.Width - float32(cameraWidth)), float32(bg.SpaceComponent.Height - float32(cameraHeight))}

}

func (*DefaultScene) Type() string { return "Game" }

func main() {
	opts := engo.RunOptions{
		Title:          "Background Image Demo",
		Width:          cameraWidth,
		Height:         cameraHeight,
		StandardInputs: true,
	}

	engo.Run(opts, &DefaultScene{})
}
