package config

import (
	"SmartRobot/internal/util"
	"github.com/spf13/viper"
)

type User struct {
	Token    string
	Channels []string
}

type FoolChatConfig struct {
	Name           string
	Users          []*User
	Random         bool // 是否随机抽取消息
	Content        []string
	Type           string // custom 自定义消息 || history 历史消息
	SleepSecondMin int64  // 最小睡眠秒数
	SleepSecondMax int64  // 最大睡眠秒数
}

type PerformChatConfig struct {
	Name                string
	Content             []string
	Tokens              []string
	Channels            []string
	TalkSleepSecond     int64
	NextTalkSleepSecond int64
}

type LuckUser struct {
	Token    string
	Channels []string
}

type LuckConfig struct {
	Name  string
	Users []*LuckUser
}

type RuleConfig struct {
	Keyword  string // 关键词过滤
	Interval int64  // 间隔发送时间
	On       bool
}

var FoolChatConfigIns *FoolChatConfig
var PerformChatConfigInc *PerformChatConfig
var LuckConfigInc *LuckConfig
var RuleConfigInc *RuleConfig

func Setup() {

	configMap := make(map[string]any)
	configMap["./config/foolchat.json"] = &FoolChatConfigIns
	configMap["./config/performchat.json"] = &PerformChatConfigInc
	configMap["./config/luck.json"] = &LuckConfigInc
	configMap["./config/rule.json"] = &RuleConfigInc

	for key, inc := range configMap {
		v := viper.New()
		// 路径必须要写相对路径,相对于项目的路径
		v.SetConfigFile(key)

		if err := v.ReadInConfig(); err != nil {
			util.Panic("读取 "+key+"  配置文件失败:", err)
		}

		// 映射到结构体
		if err := v.Unmarshal(&inc); err != nil {
			util.Panic("序列化 "+key+" 配置文件失败:", err)
		}
	}

	//// 监听配置文件变化
	//v.WatchConfig()
	//v.OnConfigChange(func(in fsnotify.Event) {
	//	fmt.Printf("配置文件:%s发生改变\n", in.Name)
	//	if err := v.ReadInConfig(); err != nil {
	//		panic(err)
	//	}
	//	if err := v.Unmarshal(&s); err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(s)
	//})
	//// 睡眠30秒让函数不要马上结束;给修改文件留出时间
	//time.Sleep(30 * time.Second)
}
