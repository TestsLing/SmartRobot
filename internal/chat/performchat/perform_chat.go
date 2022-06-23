package performchat

import (
	"SmartRobot/config"
	"SmartRobot/internal/message"
	"SmartRobot/internal/util"
	"time"
)

type PerformChat struct {
	Tokens   []string
	Channels []string
	Content  []string
}

func NewPerformChat(tokens []string, channels []string, content []string) *PerformChat {
	return &PerformChat{Tokens: tokens, Channels: channels, Content: content}
}

func (d *PerformChat) Run() {
	m := &message.CustomMessage{
		Random:              false,
		ChannelContentIndex: make(map[string]int),
		Content:             d.Content,
	}

	for _, channel := range d.Channels {
		go func(channel string) {
			for {
				for _, token := range d.Tokens {
					text := m.GetContent(channel)
					util.SendChannel(token, channel, text)

					// 回复间隔时间
					time.Sleep(time.Second * time.Duration(config.PerformChatConfigInc.TalkSleepSecond))
				}

				// 下一轮对话时间
				sec := config.PerformChatConfigInc.NextTalkSleepSecond
				time.Sleep(time.Second * time.Duration(sec))
			}

		}(channel)
	}
}
