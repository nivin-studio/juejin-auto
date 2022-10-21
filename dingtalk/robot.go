package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

type Robot struct {
	Webhook string
	Secret  string
}

func NewRobot() *Robot {
	return &Robot{
		Webhook: "",
		Secret:  "",
	}
}

func (rb *Robot) sign(t int64, secret string) string {
	strHash := fmt.Sprintf("%d\n%s", t, secret)
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(strHash))
	data := hmac256.Sum(nil)

	return base64.StdEncoding.EncodeToString(data)
}

func (rb *Robot) GetUrl() (string, error) {
	if rb.Webhook == "" {
		return "", errors.New("Robot 'url' is not set!")
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

func (rb *Robot) SetWebhook(webhook string) *Robot {
	rb.Webhook = webhook
	return rb
}

func (rb *Robot) SetSecret(secret string) *Robot {
	rb.Secret = secret
	return rb
}

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
		return fmt.Errorf("send message err: %s, message: %s", err, m.Marshaler())
	}

	code := jsoniter.Get(resp.Body(), "errcode").ToInt()
	if code != 0 {
		return fmt.Errorf("send message err: %s, message: %s", resp.Body(), m.Marshaler())
	}

	return nil
}
