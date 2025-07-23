package dummy

import (
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
)

type Block struct {
	index int
	hash  uint64
	state map[string]any
}

func (b Block) BreakInfo() block.BreakInfo {
	bl := allBlocks[b.index]

	return block.BreakInfo{
		Hardness:        bl.Hardness,
		Harvestable:     stumpTool,
		Effective:       stumpTool,
		Drops:           stumpDrops,
		BlastResistance: bl.Hardness,
	}
}

func (b Block) CanDisplace(lq world.Liquid) bool {
	bl := allBlocks[b.index]

	if !bl.CanContainLiquidSource {
		return false
	}

	w, ok := lq.(block.Water)
	return ok && w.Depth == 8 && !w.Falling
}

func (b Block) SideClosed(pos, side cube.Pos, tx *world.Tx) bool {
	return false
}

func (b Block) FlammabilityInfo() block.FlammabilityInfo {
	bl := allBlocks[b.index]
	return block.FlammabilityInfo{
		Encouragement: bl.FlameOdds,
		Flammability:  bl.FlameOdds,
		LavaFlammable: bl.FlameOdds > 0,
	}
}

func (b Block) LightDiffusionLevel() uint8 {
	bl := allBlocks[b.index]
	return uint8(bl.LightDampening)
}

func (b Block) LightEmissionLevel() uint8 {
	bl := allBlocks[b.index]
	return uint8(bl.LightEmission)
}

func (b Block) Friction() float64 {
	bl := allBlocks[b.index]
	return bl.BlockFriction
}

func (b Block) EncodeBlock() (string, map[string]any) {
	bl := allBlocks[b.index]
	return bl.Name, b.state
}

func (b Block) Hash() (uint64, uint64) {
	return b.hash, 0
}

func (b Block) Model() world.BlockModel {
	bl := allBlocks[b.index]
	return bl.Model()
}
