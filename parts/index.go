package storygame

import (
	"github.com/skilstak/storyeng-go"
	"github.com/skilstak/storygame-sample-go/parts/name"
	"github.com/skilstak/storygame-sample-go/parts/welcome"
)

var Parts = storyeng.Parts{
	"name":    name.Part,
	"welcome": welcome.Part,
}
