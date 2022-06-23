package message

import (
	"SmartRobot/config"
	"SmartRobot/internal/util"
	"fmt"
	"time"
)

type HistoryMessage struct {
	Users []*config.User
}

func NewHistoryMessage(users []*config.User) *HistoryMessage {
	return &HistoryMessage{Users: users}
}

func (h *HistoryMessage) Send() {
	for _, user := range h.Users {
		for _, channel := range user.Channels {
			// 检测规则
			if config.RuleConfigInc.On == true {
				rule := NewRule(channel, user.Token)
				time.Sleep(time.Duration(rule.KeyWordRule()))
				fmt.Println("[发送历史消息]沉默完成")
			}
			
			util.SendChannel(
				user.Token,
				channel,
				h.GetContent(user.Token, channel),
			)
		}
	}

}

func (h *HistoryMessage) GetContent(token string, channel string) string {
	message := util.HistoryMessage(token, channel)

	if len(message) <= 0 {
		customMessage := NewCustomMessage(h.Users, true, config.FoolChatConfigIns.Content)
		cm := customMessage.GetContent(channel)
		util.ErrorF("用户:%s, 获取历史消息, 已使用默认随机消息，发送通道为: %s, 消息内容为: %s", token, channel, cm)
		return cm
	}

	var contents []string

	for _, m := range message {
		contents = append(contents, m.Content)
	}

	x := util.GetRandom(int64(len(contents) - 1))
	return contents[x]
}
