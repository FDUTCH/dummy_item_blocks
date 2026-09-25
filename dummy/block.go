package dummy

import (
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/block/model"
	"github.com/df-mc/dragonfly/server/world"
)

type Block struct {
	name  string
	hash  uint64
	state map[string]any
}

func (b Block) BreakInfo() block.BreakInfo {
	bl := allBlocks[b.name]

	return block.BreakInfo{
		Hardness:        bl.Hardness,
		Harvestable:     stumpTool,
		Effective:       stumpTool,
		Drops:           stumpDrops,
		BlastResistance: bl.Hardness,
	}
}

func (b Block) SideClosed(pos, side cube.Pos, tx *world.Tx) bool {
	return false
}

func (b Block) FlammabilityInfo() block.FlammabilityInfo {
	bl := allBlocks[b.name]
	return block.FlammabilityInfo{
		Encouragement: bl.FlameEncouragement,
		Flammability:  bl.Flammability,
		LavaFlammable: bl.FlameEncouragement > 0,
	}
}

func (b Block) Friction() float64 {
	bl, ok := allBlocks[b.name]
	if !ok {
		return 0.6
	}
	return bl.BlockFriction
}

func (b Block) EncodeBlock() (string, map[string]any) {
	return b.name, b.state
}

func (b Block) Hash() (uint64, uint64) {
	return b.hash, 0
}

func (b Block) Model() world.BlockModel {
	return model.Solid{}
}
