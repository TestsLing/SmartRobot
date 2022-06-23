package util

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	url2 "net/url"
	"strings"
	"time"
)

const (
	Domain            = "https://discord.com/"
	SendMessageUrl    = Domain + "api/v9/channels/%s/messages"
	MessageHistoryUrl = Domain + "api/v9/channels/%s/messages?limit=100"
	LuckUrl           = Domain + "api/v9/channels/%s/messages/%s/reactions/%s/%s?location=Message"
	Nonce             = "82329451214%d33232234"
	ContentType       = "application/json"
	UserAgent         = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36"
)

const (
	HttpProxy  = "http://127.0.0.1:10809"
	SocksProxy = "socks5://127.0.0.1:10808"
)

func GetNonce() string {
	return fmt.Sprintf(Nonce, GetRandom(1000))
}

func GetRandom(num int64) int64 {
	bigInt, _ := rand.Int(rand.Reader, big.NewInt(num))
	return bigInt.Int64()
}

func RangeRand(min, max int64) int64 {
	if min > max {
		panic("the min is greater than max!")
	}

	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))

		return result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}

func SendChannel(token string, channel string, content string) {

	data := make(map[string]any)
	data["content"] = content
	data["nonce"] = GetNonce()
	data["ttl"] = false
	InfoF("发送通道消息 用户是:%s, 发送的通道是:%s, 发送的内容是:%s", token, channel, content)

	url := fmt.Sprintf(SendMessageUrl, channel)

	bodyData, err := json.Marshal(data)
	if err != nil {
		Fatal("发送通道消息请求体构造失败 err:", err)
		return
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: func(_ *http.Request) (*url2.URL, error) {
				return url2.Parse(HttpProxy)
			},
		},
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyData))

	if err != nil {
		Fatal("发送通道消息请求对象构造失败 err:", err)
		return
	}

	req.Header.Add("Content-Type", ContentType)
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Authorization", token)

	rep, err := client.Do(req)

	if err != nil {
		Fatal("发送消息失败 err:", err)
		return
	}

	if rep.StatusCode != http.StatusOK {
		Fatal("发送消息响应码不为200 err:", err)
		return
	}

	b, err := ioutil.ReadAll(rep.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			Fatal("发送消息响应体关闭失败 err:", err)

		}
	}(rep.Body)

	Info("请求 SendChannel 成功，响应数据为:", string(b))
}

type Message struct {
	ID        string
	Content   string
	Channel   string
	Timestamp time.Time
}

func HistoryMessage(token string, channel string) []Message {
	var m []Message

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: func(_ *http.Request) (*url2.URL, error) {
				return url2.Parse(HttpProxy)
			},
		},
	}

	url := fmt.Sprintf(MessageHistoryUrl, channel)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		Fatal("构造 HistoryMessage Get 请求失败, err:", err)
		return m
	}

	req.Header.Add("Content-Type", ContentType)
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Authorization", token)

	rep, err := client.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			Fatal("关闭 HistoryMessage Get 请求 Body 失败, err:", err)
		}
	}(rep.Body)

	if rep.StatusCode != http.StatusOK {
		Fatal("HistoryMessage Get 请求响应码不为 200, err:", err)
		return m
	}

	b, err := ioutil.ReadAll(rep.Body)

	err = json.NewDecoder(strings.NewReader(string(b))).Decode(&m)
	if err != nil {
		Fatal("HistoryMessage Get 响应数据序列化失败, err:", err)
		return m
	}

	Info("请求 HistoryMessage 成功，响应数据为:", m)

	return m
}

type Emoji struct {
	ID   string
	Name string
}

type ReactionItem struct {
	Emoji Emoji
	Count int
	Me    bool
}

type LuckMessage struct {
	ID        string
	ChannelId string
	Reactions []*ReactionItem
}

func HistoryLuckMessage(token string, channel string) []LuckMessage {
	var m []LuckMessage

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: func(_ *http.Request) (*url2.URL, error) {
				return url2.Parse(HttpProxy)
			},
		},
	}

	url := fmt.Sprintf(MessageHistoryUrl, channel)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		Fatal("构造 HistoryLuckMessage Get 请求失败, err:", err)
		return m
	}

	req.Header.Add("Content-Type", ContentType)
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Authorization", token)

	rep, err := client.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			Fatal("关闭 HistoryLuckMessage Get 请求 Body 失败, err:", err)
		}
	}(rep.Body)

	if rep.StatusCode != http.StatusOK {
		Fatal("HistoryLuckMessage Get 请求响应码不为 200, err:", err)
		return m
	}

	b, err := ioutil.ReadAll(rep.Body)

	err = json.NewDecoder(strings.NewReader(string(b))).Decode(&m)
	if err != nil {
		Fatal("HistoryLuckMessage Get 响应数据序列化失败, err:", err)
		return m
	}

	Info("请求 HistoryLuckMessage 成功，响应数据为:", m)

	return m
}

func JoinLuck(token string, channel string, messageId string, emoji string) {

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: func(_ *http.Request) (*url2.URL, error) {
				return url2.Parse(HttpProxy)
			},
		},
	}

	url := fmt.Sprintf(LuckUrl, channel, messageId, emoji, "@me")
	req, err := http.NewRequest("PUT", url, nil)

	if err != nil {
		Fatal("构造 HistoryLuckMessage Put 请求失败, err:", err)
	}

	req.Header.Add("Content-Type", ContentType)
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Authorization", token)

	rep, err := client.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			Fatal("关闭 HistoryLuckMessage Put 请求 Body 失败, err:", err)
		}
	}(rep.Body)

	if rep.StatusCode != http.StatusNoContent {
		Fatal("HistoryLuckMessage Put 请求响应码不为 204, err:", err)
	}

	b, err := ioutil.ReadAll(rep.Body)

	if err != nil {
		Fatal("HistoryLuckMessage Put 响应数据序列化失败, err:", err)
	}

	Info("请求 HistoryLuckMessage 成功，响应数据为:", string(b))

}
