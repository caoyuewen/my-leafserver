package main

import (
	"github.com/name5566/leaf"
	"my-leafserver/game"
	"my-leafserver/gate"
	"my-leafserver/login"
	_ "my-leafserver/models"
)

func main() {
	//lconf.LogLevel = conf.Server.LogLevel
	//lconf.LogPath = conf.Server.LogPath
	//lconf.ConsolePort = conf.Server.ConsolePort
	//lconf.ProfilePath = conf.Server.ProfilePath
	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
