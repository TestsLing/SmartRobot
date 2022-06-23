package main

import (
	"SmartRobot/config"
	_ "SmartRobot/config"
	"SmartRobot/internal/chat/foolchat"
	"SmartRobot/internal/util"
)

func main() {
	config.Setup()
	util.SetupLog()

	chat := foolchat.NewFoolChat(config.FoolChatConfigIns.Type, config.FoolChatConfigIns.Users, config.FoolChatConfigIns.Random, config.FoolChatConfigIns.Content)
	go chat.Run()

	//performChat := performchat.NewPerformChat(config.PerformChatConfigInc.Tokens, config.PerformChatConfigInc.Channels, config.PerformChatConfigInc.Content)
	//go performChat.Run()
	//
	//newLuck := luck.NewLuck(config.LuckConfigInc.Users)
	//newLuck.Run()

	select {}
}
