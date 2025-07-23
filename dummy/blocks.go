package dummy

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/world"
	"log/slog"
	"reflect"
	_ "unsafe"
)

var (
	unregisteredBlocks = make(map[string]world.Block)

	//go:linkname blocks github.com/df-mc/dragonfly/server/world.blocks
	blocks []world.Block

	allBlocks []blockData

	EnabledLogging bool
)

func isRegistered(bl world.Block) bool {
	return "unknownBlock" != reflect.ValueOf(bl).Type().Name()
}

func Register() {
	parseBlockData()
	paseItemData()

	var (
		registeredItems  int
		registeredBlocks int
	)

	for index, b := range blocks {
		if isRegistered(b) {
			continue
		}
		name, state := b.EncodeBlock()
		bl := Block{
			index: index,
			hash:  block.NextHash(),
			state: state,
		}
		itemName, isItem := blockToItem[name]
		if _, registered := world.ItemByName(itemName, 0); !registered && isItem {
			it := ItemBlock{
				Block:    bl,
				itemName: itemName,
			}
			world.RegisterItem(it)
			world.RegisterBlock(it)
			registeredBlocks++
			registeredItems++
			continue
		}
		world.RegisterBlock(bl)
		registeredBlocks++
	}
	if EnabledLogging {
		slog.Info(fmt.Sprintf("there were registered %d new items and %d new blocks", registeredItems, registeredBlocks))
	}
}
