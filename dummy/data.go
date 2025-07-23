package dummy

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/block/model"
	"github.com/df-mc/dragonfly/server/world"
	"unsafe"
)

var (
	//go:embed block_states_raw.json
	encodedBlockData []byte //https://github.com/AllayMC/Allay/blob/master/data/resources/unpacked/block_states_raw.json
)

type blockData struct {
	BurnOdds                    int          `json:"burnOdds"`
	CanContainLiquidSource      bool         `json:"canContainLiquidSource"`
	CollisionShape              [][6]float64 `json:"collisionShape"`
	ExplosionResistance         float64      `json:"explosionResistance"`
	FlameOdds                   int          `json:"flameOdds"`
	BlockFriction               float64      `json:"friction"`
	Hardness                    float64      `json:"hardness"`
	IsSolid                     bool         `json:"isSolid"`
	LightDampening              int          `json:"lightDampening"`
	LightEmission               int          `json:"lightEmission"`
	LiquidReactionOnTouch       string       `json:"liquidReactionOnTouch"`
	Name                        string       `json:"name"`
	RequiresCorrectToolForDrops bool         `json:"requiresCorrectToolForDrops"`
}

func (s blockData) CanDisplace(b world.Liquid) bool {
	if !s.CanContainLiquidSource {
		return false
	}

	w, ok := b.(block.Water)
	return ok && w.Depth == 8 && !w.Falling
}

func (s blockData) SideClosed(pos, side cube.Pos, tx *world.Tx) bool {
	return false
}

func (s blockData) Model() world.BlockModel {
	if s.IsSolid {
		return model.Solid{}
	}
	return Model{boxes: *(*[]cube.BBox)(unsafe.Pointer(&s.CollisionShape))}
}

func parseBlockData() {
	err := json.NewDecoder(bytes.NewReader(encodedBlockData)).Decode(&allBlocks)
	if err != nil {
		panic(err)
	}
}
