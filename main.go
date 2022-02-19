package main

import (
	"embed"
	"github.com/jjanvier/tdd/cmd"
	"github.com/jjanvier/tdd/handler"
)

//go:embed assets/*
var assets embed.FS

func main() {
	handler.Assets = assets
	cmd.Execute()
}
