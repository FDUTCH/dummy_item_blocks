package dummy

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"github.com/df-mc/dragonfly/server/block"
)

type ItemBlock struct {
	Block
	itemName string
}

var (
	//go:embed block_id_to_item_id_map.json
	encodedItemData []byte //https://github.com/pmmp/BedrockData/blob/master/block_id_to_item_id_map.json

	blockToItem = make(map[string]string)
)

func (b ItemBlock) EncodeItem() (name string, meta int16) {
	return b.itemName, 0
}

func (b ItemBlock) BreakInfo() block.BreakInfo {
	bl := allBlocks[b.index]

	return block.BreakInfo{
		Hardness:        bl.Hardness,
		Harvestable:     stumpTool,
		Effective:       stumpTool,
		Drops:           oneOf(b),
		BlastResistance: bl.Hardness,
	}
}

func paseItemData() {
	err := json.NewDecoder(bytes.NewReader(encodedItemData)).Decode(&blockToItem)
	if err != nil {
		panic(err)
	}
}
