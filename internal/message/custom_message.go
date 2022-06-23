package message

import (
	"SmartRobot/config"
	"SmartRobot/internal/util"
	"fmt"
	"time"
)

type CustomMessage struct {
	Users               []*config.User
	Random              bool
	ChannelContentIndex map[string]int
	Content             []string
}

func NewCustomMessage(users []*config.User, random bool, content []string) *CustomMessage {
	return &CustomMessage{Users: users, Random: random, ChannelContentIndex: make(map[string]int), Content: content}
}

func (c *CustomMessage) Send() {
	for _, user := range c.Users {
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
				c.GetContent(channel),
			)
		}
	}

}

func (c *CustomMessage) GetContent(channel string) string {
	content := c.Content
	random := c.Random

	if random == true {
		x := util.GetRandom(int64(len(content) - 1))
		return content[x]
	}

	index := c.ChannelContentIndex[channel]
	str := content[index]

	if index >= len(content)-1 {
		index = 0
	} else {
		index++
	}

	c.ChannelContentIndex[channel] = index

	return str
}
