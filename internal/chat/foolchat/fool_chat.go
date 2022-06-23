package foolchat

import (
	"SmartRobot/config"
	message2 "SmartRobot/internal/message"
	"SmartRobot/internal/util"
	"time"
)

const (
	Custom  = "custom"
	History = "history"
)

type FoolChat struct {
	Type    string
	Users   []*config.User
	Random  bool
	Content []string
}

func NewFoolChat(t string, users []*config.User, random bool, content []string) *FoolChat {
	return &FoolChat{Type: t, Users: users, Random: random, Content: content}
}

func (f *FoolChat) Run() {
	var message message2.IMessage

	for {
		switch f.Type {
		case History:
			message = message2.NewHistoryMessage(f.Users)
			message.Send()
		case Custom:
			message = message2.NewCustomMessage(f.Users, f.Random, f.Content)
			message.Send()
		default:
			message = message2.NewCustomMessage(f.Users, f.Random, f.Content)
			message.Send()
		}
		min := config.FoolChatConfigIns.SleepSecondMin
		max := config.FoolChatConfigIns.SleepSecondMax
		sleepSecond := util.RangeRand(min, max)
		time.Sleep(time.Second * time.Duration(sleepSecond))
	}

}
