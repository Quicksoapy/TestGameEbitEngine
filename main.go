package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math"
)

// Game implements ebiten.Game interface.
type Game struct {
	circle1      *ebiten.Image
	circle2      *ebiten.Image
	circle3      *ebiten.Image
	circleCenter *ebiten.Image
	deg          int
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	if g.circleCenter == nil {
		c, _, err := ebitenutil.NewImageFromFile("Assets/CircleCenter.png")
		if err != nil {
			return err
		}
		g.circleCenter = c
	}
	g.deg++

	if g.circle1 == nil {
		c, _, err := ebitenutil.NewImageFromFile("Assets/Circle1.png")
		if err != nil {
			return err
		}
		g.circle1 = c
	}

	if g.circle2 == nil {
		c, _, err := ebitenutil.NewImageFromFile("Assets/Circle2.png")
		if err != nil {
			return err
		}
		g.circle2 = c
	}

	if g.circle3 == nil {
		c, _, err := ebitenutil.NewImageFromFile("Assets/Circle3.png")
		if err != nil {
			return err
		}
		g.circle3 = c
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	width, height := screen.Size()
	/*
		Circle1, _, err := ebitenutil.NewImageFromFile("Assets/Circle1.png")
		if err != nil {
			log.Fatal(err)
		}
		tempVarCr1 := DrawImageOptionsGenerator(*Circle1, width, height)
		screen.DrawImage(Circle1, &tempVarCr1)

		Circle2, _, err := ebitenutil.NewImageFromFile("Assets/Circle2.png")
		if err != nil {
			log.Fatal(err)
		}
		tempVarCr2 := DrawImageOptionsGenerator(*Circle2, width, height)
		screen.DrawImage(Circle2, &tempVarCr2)

		Circle3, _, err := ebitenutil.NewImageFromFile("Assets/Circle3.png")
		if err != nil {
			log.Fatal(err)
		}
		tempVarCr3 := DrawImageOptionsGenerator(*Circle3, width, height)
		screen.DrawImage(Circle3, &tempVarCr3)
	*/
	if g.circleCenter == nil {
		return
	}
	tempVarCrCenter := DrawImageOptionsGenerator(g.circleCenter, width, height) // Assumes that DrawImageOptionsGenerator takes a pointer of ebiten.Image

	w, h := float64(g.circleCenter.Bounds().Dx()), float64(g.circleCenter.Bounds().Dy())
	tempVarCrCenter.GeoM.Translate(-w/2, -h/2)
	tempVarCrCenter.GeoM.Rotate(float64(g.deg) * math.Pi / 180)
	tempVarCrCenter.GeoM.Translate(w/2, h/2)

	screen.DrawImage(g.circleCenter, &tempVarCrCenter)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("meow")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func DrawImageOptionsGenerator(Circle *ebiten.Image, ScreenWidth int, ScreenHeight int) ebiten.DrawImageOptions {
	drawimageoptions := ebiten.DrawImageOptions{}
	width, height := Circle.Size()
	drawimageoptions.GeoM.Scale(0.7, 0.7)
	drawimageoptions.GeoM.Translate(ImageTranslator(float32(ScreenWidth/2), float32(width)*0.7), ImageTranslator(float32(ScreenHeight/2), float32(height)*0.7))

	return drawimageoptions
}

func ImageTranslator(middle, size float32) float64 {
	return float64(middle - size/2)
}
