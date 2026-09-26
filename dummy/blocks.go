package dummy

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"log/slog"
	"reflect"
	_ "unsafe"

	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/creative"
	"github.com/df-mc/dragonfly/server/world"
)

var (
	allBlocks map[string]blockData

	Logging bool
)

func isRegistered(bl world.Block) bool {
	return "unknownBlock" != reflect.ValueOf(bl).Type().Name()
}

func Register(registry *world.BasicBlockRegistry) {
	eduGroup := creative.Group{
		Category: creative.ItemsCategory(),
		Name:     "edu",
		Icon:     item.NewStack(item.EnderEye{}, 1),
	}

	creative.RegisterGroup(eduGroup)

	parseBlockData()
	paseItemData()

	var (
		registeredItems  int
		registeredBlocks int
	)

	c := registry.Clone()
	c.Finalize()

	for _, b := range c.Blocks() {
		if isRegistered(b) {
			continue
		}

		name, _ := b.EncodeBlock()

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

			_ = it

			world.RegisterItem(it)
			registry.RegisterBlock(it)

			if _, has := eduContent[itemName]; has {
				creative.RegisterItem(creative.Item{
					Stack: item.NewStack(it, 1),
					Group: "edu",
				})
			}

			registeredBlocks++
			registeredItems++
			continue
		}
		registry.RegisterBlock(bl)
		registeredBlocks++
	}
	if Logging {
		slog.Info(fmt.Sprintf("there were registered %d new items and %d new blocks", registeredItems, registeredBlocks))
	}
}

var eduContent = make(map[string]struct{})

//go:embed penis.txt
var eduBlocks []byte

func init() {
	s := bufio.NewScanner(bytes.NewReader(eduBlocks))

	for s.Scan() {
		eduContent[s.Text()] = struct{}{}
	}
}
