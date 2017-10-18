package storygame

import (
	game "github.com/skilstak/storyeng-go"
	"github.com/skilstak/storygame-sample-go/parts/name"
	"github.com/skilstak/storygame-sample-go/parts/welcome"
)

var Parts = game.Parts{
	"name":    name.Part,
	"welcome": welcome.Part,
}
