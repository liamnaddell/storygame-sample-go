package main

import (
	game "github.com/skilstak/storyeng-go"
	storygame "github.com/skilstak/storygame-sample-go/parts"
)

func main() {
	story := game.NewStory(storygame.Parts)
	story.Start()
}
