package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 0.15
	poolSize = 20
)

type bullet struct {
	tex   *sdl.Texture
	x, y  float64
	angle float64

	active bool
}

// newBullet gives you new bullet
func newBullet(renderer *sdl.Renderer) (bul bullet) {
	bul.tex = textureFromBMP(renderer, "images/player_bullet.bmp")
	return bul
}

func (bul *bullet) draw(renderer *sdl.Renderer) {
	if !bul.active {
		return
	}

	// Converting bullet coordinates to top left of sprite
	x := bul.x - bulletSize/2.0
	y := bul.y - bulletSize/2.0

	renderer.Copy(bul.tex,
		&sdl.Rect{X: 0, Y: 0, W: bulletSize, H: bulletSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: bulletSize, H: bulletSize})
}

// update, updates the position of the object and set inactive to the fired bullets
func (bul *bullet) update() {
	bul.x += bulletSpeed * math.Cos(bul.angle)
	bul.y += bulletSpeed * math.Sin(bul.angle)

	if bul.x > screenWidth || bul.x < 0 || bul.y > screenHeight || bul.y < 0 {
		bul.active = false // part of Object Pool
	}
}

var bulletPool []*bullet

/* 
	The below two functions are the part of object pool
*/

// initBulletPool initialized the pool with poolSize objects
func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < poolSize; i++ {
		bul := newBullet(renderer)
		bulletPool = append(bulletPool, &bul)
	}
}


// bulletFromPool gives you inactive bullets from the bulletPool
func bulletFromPool() (*bullet, bool) {
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true
		}
	}

	return nil, false
}
