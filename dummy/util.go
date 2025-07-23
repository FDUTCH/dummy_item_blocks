package dummy

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

func oneOf(i ...world.Item) func(item.Tool, []item.Enchantment) []item.Stack {
	return func(item.Tool, []item.Enchantment) []item.Stack {
		var s []item.Stack
		for _, it := range i {
			s = append(s, item.NewStack(it, 1))
		}
		return s
	}
}

func stumpTool(_ item.Tool) bool {
	return true
}

func stumpDrops(_ item.Tool, _ []item.Enchantment) []item.Stack {
	return nil
}
