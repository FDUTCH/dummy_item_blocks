package dummy

import (
	_ "embed"
	"encoding/json"
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/model"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
	"reflect"
	_ "unsafe"
)

func Register() {
	fillUnregistered()
	for name, bl := range unregisteredBlocks {
		data, ok := blocksData[name]
		if !ok {
			continue
		}
		bl = data.Block(bl)
		world.RegisterBlock(bl)
		if it, ok := bl.(world.Item); ok {
			name, id := it.EncodeItem()
			_, ok = world.ItemByName(name, id)
			if !ok {
				world.RegisterItem(it)
			}
		}
	}
}

type BlockItem struct {
	Block
	itemName string
}
type Block struct {
	hash             uint64
	name             string
	flammabilityInfo block.FlammabilityInfo
	breakInfo        block.BreakInfo
	fiction          float64
	lightDiffusion   uint8
	lightEmission    uint8
	state            map[string]any
}

func (b BlockItem) EncodeItem() (name string, meta int16) {
	return b.itemName, 0
}

func (b Block) LightEmissionLevel() uint8 {
	return b.lightEmission
}

func (b Block) LightDiffusionLevel() uint8 {
	return b.lightDiffusion
}

func (b Block) Friction() float64 {
	return b.fiction
}

func (b Block) BreakInfo() block.BreakInfo {
	return b.breakInfo
}

func (b Block) FlammabilityInfo() block.FlammabilityInfo {
	return b.flammabilityInfo
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

type Data struct {
	BlastResistance    float64 `json:"blastResistance"`
	Brightness         float64 `json:"brightness"`
	FlameEncouragement int     `json:"flameEncouragement"`
	Flammability       int     `json:"flammability"`
	Friction           float64 `json:"friction"`
	Hardness           float64 `json:"hardness"`
	Opacity            float64 `json:"opacity"`
}

func (d Data) Block(unregistered world.Block) world.Block {
	name, state := unregistered.EncodeBlock()
	itemName, has := blockToItemName[name]
	bl := Block{
		hash:             block.NextHash(),
		name:             name,
		flammabilityInfo: block.FlammabilityInfo{},
		breakInfo: block.BreakInfo{
			Hardness:        d.Hardness,
			BlastResistance: d.BlastResistance,
			Harvestable:     stumpTool,
			Effective:       stumpTool,
		},
		fiction:        d.Friction,
		lightDiffusion: uint8(15 * float64(d.Opacity)),
		lightEmission:  uint8(d.Brightness),
		state:          state,
	}
	if has {
		blItem := BlockItem{
			Block:    bl,
			itemName: itemName,
		}
		blItem.breakInfo.Drops = oneOf(blItem)
		return blItem
	}
	bl.breakInfo.Drops = stumpDrops
	return bl
}

var (
	//go:embed block_id_to_item_id_map.json
	blockIdToItemId []byte
	blockToItemName map[string]string

	//go:embed block_properties_table.json
	blockProperties []byte
	blocksData      map[string]Data

	unregisteredBlocks = make(map[string]world.Block)
)

func stumpTool(t item.Tool) bool {
	return true
}

func stumpDrops(t item.Tool, enchantments []item.Enchantment) []item.Stack {
	return nil
}

func init() {
	err := json.Unmarshal(blockIdToItemId, &blockToItemName)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(blockProperties, &blocksData)
	if err != nil {
		panic(err)
	}
}

//go:linkname blocks github.com/df-mc/dragonfly/server/world.blocks
var blocks []world.Block

func fillUnregistered() {
	for _, bl := range blocks {
		if "unknownBlock" == reflect.ValueOf(bl).Type().Name() {
			str, _ := bl.EncodeBlock()
			_, has := unregisteredBlocks[str]
			if has {
				continue
			}
			unregisteredBlocks[str] = bl
		}
	}
}

func oneOf(i ...world.Item) func(item.Tool, []item.Enchantment) []item.Stack {
	return func(item.Tool, []item.Enchantment) []item.Stack {
		var s []item.Stack
		for _, it := range i {
			s = append(s, item.NewStack(it, 1))
		}
		return s
	}
}
