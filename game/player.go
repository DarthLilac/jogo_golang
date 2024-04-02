package game

import (

	"jogo_golang/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image *ebiten.Image
	position Vector
}

func NewPlayer() *Player {
	image := assets.PlayerSprite

	position := Vector{
		X: (screenWidth / 2),
		Y: 500,
	}
	
	return &Player{
		image: image,
		position: position,
	}
}

func(p *Player) Update() {

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.image, op)

}