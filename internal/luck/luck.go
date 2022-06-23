package luck

import (
	"SmartRobot/config"
	"SmartRobot/internal/util"
	"time"
)

type Luck struct {
	Users []*config.LuckUser
}

func NewLuck(users []*config.LuckUser) *Luck {
	return &Luck{Users: users}
}

func (l *Luck) Run() {
	for _, user := range l.Users {
		for _, channel := range user.Channels {
			// 获取历史消息
			luckMessage := util.HistoryLuckMessage(user.Token, channel)

			for _, message := range luckMessage {
				if len(message.Reactions) <= 0 {
					continue
				}
				for _, reaction := range message.Reactions {
					if reaction.Me == true {
						continue
					}
					// 没点的全他妈点一遍
					util.JoinLuck(user.Token, channel, message.ID, reaction.Emoji.Name)
					time.Sleep(time.Second * 3)
				}
			}

		}
		util.Info("用户: ", user.Token, "参与抽奖完成")
	}

	util.Info("全部用户参与抽奖完成！！！")

}
