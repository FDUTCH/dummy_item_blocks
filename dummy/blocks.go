package dummy

import (
	"fmt"
	"log/slog"
	"reflect"
	_ "unsafe"

	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/world"
)

var (
	allBlocks map[string]blockData

	Logging bool
)

func isRegistered(bl world.Block) bool {
	return "unknownBlock" != reflect.ValueOf(bl).Type().Name()
}

func Register(registry world.BlockRegistry) {
	parseBlockData()
	paseItemData()

	var (
		registeredItems  int
		registeredBlocks int
	)

	for _, b := range registry.Blocks() {
		if isRegistered(b) {
			continue
		}
		name, state := b.EncodeBlock()
		bl := Block{
			name:  name,
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
	if Logging {
		slog.Info(fmt.Sprintf("there were registered %d new items and %d new blocks", registeredItems, registeredBlocks))
	}
}
