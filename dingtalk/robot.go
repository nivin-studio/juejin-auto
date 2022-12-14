package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

type Robot struct {
	Webhook string
	Secret  string
}

// 创建钉钉机器人实例
func NewRobot() *Robot {
	return &Robot{
		Webhook: "",
		Secret:  "",
	}
}

// 签名
func (rb *Robot) sign(t int64, secret string) string {
	strHash := fmt.Sprintf("%d\n%s", t, secret)
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(strHash))
	data := hmac256.Sum(nil)

	return base64.StdEncoding.EncodeToString(data)
}

// 获取请求地址
func (rb *Robot) GetUrl() (string, error) {
	if rb.Webhook == "" {
		return "", fmt.Errorf("钉钉机器人Webhook未设置")
	}

	if rb.Secret == "" {
		return rb.Webhook, nil
	} else {
		value := url.Values{}
		times := time.Now().UnixNano() / int64(time.Millisecond)

		value.Set("timestamp", fmt.Sprintf("%d", times))
		value.Set("sign", rb.sign(times, rb.Secret))

		return rb.Webhook + "&" + value.Encode(), nil
	}
}

// 设置webhook地址
func (rb *Robot) SetWebhook(webhook string) *Robot {
	rb.Webhook = webhook
	return rb
}

// 设置签名密钥
func (rb *Robot) SetSecret(secret string) *Robot {
	rb.Secret = secret
	return rb
}

// 发送消息
func (rb *Robot) SendMessage(m *Message) error {
	url, err := rb.GetUrl()
	if err != nil {
		return err
	}

	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(m.Marshaler()).
		Post(url)

	if err != nil {
		return fmt.Errorf("消息发送失败: %s, 消息: %s", err, m.Marshaler())
	}

	log.Println("消息发送请求结果:", string(resp.Body()))

	code := jsoniter.Get(resp.Body(), "errcode").ToInt()
	if code != 0 {
		return fmt.Errorf("消息发送失败: %s, 消息: %s", resp.Body(), m.Marshaler())
	}

	return nil
}
