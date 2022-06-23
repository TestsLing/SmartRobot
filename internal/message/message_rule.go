package message

import (
	"SmartRobot/config"
	"SmartRobot/internal/util"
	"fmt"
	"sort"
	"strings"
	"time"
)

type Rule struct {
	Message []util.Message
	Channel string
	Token   string
}

func NewRule(channel string, token string) *Rule {
	message := util.HistoryMessage(token, channel)
	return &Rule{Channel: channel, Message: message}
}

func (r *Rule) KeyWordRule() int64 {
	keyword := config.RuleConfigInc.Keyword

	ms := r.Message

	// 按照时间进行排序
	sort.Slice(ms, func(i, j int) bool {
		return ms[i].Timestamp.Unix() > ms[j].Timestamp.Unix()
	})

	var keywordMessage *util.Message

	for i := 0; i < len(ms); i++ {
		fmt.Println(ms[i])
		if strings.Contains(ms[i].Content, keyword) {
			keywordMessage = &ms[i]
			fmt.Println("[检测到关键字]消息为:", keywordMessage)
			break
		}
	}

	if keywordMessage == nil {
		return 0
	}

	// 找到了包含关键字的消息 判断时间是否超过了间隔时间
	now := time.Now()

	interval := now.Unix() - keywordMessage.Timestamp.Unix()

	if interval >= config.RuleConfigInc.Interval {
		fmt.Println("不用沉默")
		return 0
	}

	fmt.Println("直接沉默")
	return interval
}
