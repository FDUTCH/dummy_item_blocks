package dummy

import (
	"bytes"
	_ "embed"
	"encoding/json"
)

var (
	//go:embed block_properties_table.json
	encodedBlockData []byte //https://github.com/axolotl-pm/BedrockData/blob/master/block_properties_table.json
)

type blockData struct {
	BlastResistance    float64 `json:"blastResistance"`
	Brightness         float64 `json:"brightness"`
	FlameEncouragement int     `json:"flameEncouragement"`
	Flammability       int     `json:"flammability"`
	BlockFriction      float64 `json:"friction"`
	Hardness           float64 `json:"hardness"`
	Opacity            float64 `json:"opacity"`
}

//func (s blockData) CanDisplace(b world.Liquid) bool {
//	if !s.CanContainLiquidSource {
//		return false
//	}
//
//	w, ok := b.(block.Water)
//	return ok && w.Depth == 8 && !w.Falling
//}

//func (s blockData) SideClosed(pos, side cube.Pos, tx *world.Tx) bool {
//	return false
//}

//func (s blockData) Model() world.BlockModel {
//	if s.IsSolid {
//		return model.Solid{}
//	}
//	return Model{boxes: *(*[]cube.BBox)(unsafe.Pointer(&s.CollisionShape))}
//}

func parseBlockData() {
	err := json.NewDecoder(bytes.NewReader(encodedBlockData)).Decode(&allBlocks)
	if err != nil {
		panic(err)
	}
}
