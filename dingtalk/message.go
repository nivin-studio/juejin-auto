package dingtalk

import (
	"encoding/json"
)

const (
	MsgTypeText     = "text"
	MsgTypeLink     = "link"
	MsgTypeMarkdown = "markdown"
)

type TextMessage struct {
	Content string `json:"content,omitempty"`
}

type LinkMessage struct {
	Title      string `json:"title,omitempty"`
	Text       string `json:"text,omitempty"`
	PicUrl     string `json:"picUrl,omitempty"`
	MessageUrl string `json:"messageUrl,omitempty"`
}

type MarkdownMessage struct {
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
}

type AtMessage struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}

type Message struct {
	MsgType  string           `json:"msgtype,omitempty"`
	Text     *TextMessage     `json:"text,omitempty"`
	Link     *LinkMessage     `json:"link,omitempty"`
	Markdown *MarkdownMessage `json:"markdown,omitempty"`
	At       *AtMessage       `json:"at,omitempty"`
}

// 创建文本消息
func NewTextMessage(content string) *Message {
	message := &Message{
		MsgType: MsgTypeText,
		Text: &TextMessage{
			Content: content,
		},
	}

	return message
}

// 创建链接消息
func NewLinkMessage(title string, text string, picUrl string, messageUrl string) *Message {
	message := &Message{
		MsgType: MsgTypeLink,
		Link: &LinkMessage{
			Title:      title,
			Text:       text,
			PicUrl:     picUrl,
			MessageUrl: messageUrl,
		},
	}

	return message
}

// 创建markdown消息
func NewMarkdownMessage(title string, text string) *Message {
	message := &Message{
		MsgType: MsgTypeMarkdown,
		Markdown: &MarkdownMessage{
			Title: title,
			Text:  text,
		},
	}

	return message
}

// @所有人
func (m *Message) AtAll() *Message {
	m.At = &AtMessage{
		IsAtAll: true,
	}

	return m
}

// @指定用户
func (m *Message) AtMobiles(mobiles ...string) *Message {
	m.At = &AtMessage{}

	m.At.AtMobiles = append(m.At.AtMobiles, mobiles...)

	return m
}

// 消息序列号
func (m *Message) Marshaler() []byte {
	b, _ := json.Marshal(m)
	return b
}
