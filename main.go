package main

import (
	"github.com/skilstak/storyeng-go"
	"github.com/skilstak/storygame-sample-go/parts"
)

func main() {
	parts := storygame.Parts
	story := storyeng.NewStory(parts)
	story.Start()
}
