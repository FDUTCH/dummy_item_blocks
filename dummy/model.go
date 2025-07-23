package dummy

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
)

type Model struct {
	boxes []cube.BBox
}

func (m Model) BBox(_ cube.Pos, _ world.BlockSource) []cube.BBox {
	return m.boxes
}

func (m Model) FaceSolid(_ cube.Pos, face cube.Face, _ world.BlockSource) bool {
	for _, box := range m.boxes {
		switch face {
		case cube.FaceUp:
			if box.Max().Y() == 0.5 {
				return true
			}
		case cube.FaceDown:
			if box.Min().Y() == -0.5 {
				return true
			}
		case cube.FaceSouth:
			if box.Max().Z() == 0.5 {
				return true
			}
		case cube.FaceNorth:
			if box.Min().Z() == -0.5 {
				return true
			}
		case cube.FaceEast:
			if box.Max().X() == 0.5 {
				return true
			}
		case cube.FaceWest:
			if box.Min().X() == -0.5 {
				return true
			}
		}
	}
	return false
}
