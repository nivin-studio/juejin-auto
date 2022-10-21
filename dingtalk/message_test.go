package dingtalk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTextMessage(t *testing.T) {
	m1 := NewTextMessage("content")

	assert.Equal(t, `{"msgtype":"text","text":{"content":"content"}}`, string(m1.Marshaler()))

	m2 := NewTextMessage("content").AtAll()

	assert.Equal(t, `{"msgtype":"text","text":{"content":"content"},"at":{"isAtAll":true}}`, string(m2.Marshaler()))

	m3 := NewTextMessage("content").AtMobiles("152111111111")

	assert.Equal(t, `{"msgtype":"text","text":{"content":"content"},"at":{"atMobiles":["152111111111"]}}`, string(m3.Marshaler()))

	m4 := NewTextMessage("content").AtMobiles("152111111111", "152111111112")

	assert.Equal(t, `{"msgtype":"text","text":{"content":"content"},"at":{"atMobiles":["152111111111","152111111112"]}}`, string(m4.Marshaler()))
}

func TestNewLinkMessage(t *testing.T) {
	m := NewLinkMessage("title", "text", "picUrl", "messageUrl")

	assert.Equal(t, `{"msgtype":"link","link":{"title":"title","text":"text","picUrl":"picUrl","messageUrl":"messageUrl"}}`, string(m.Marshaler()))
}

func TestNewMarkdownMessage(t *testing.T) {
	m := NewMarkdownMessage("title", "text")

	assert.Equal(t, `{"msgtype":"markdown","markdown":{"title":"title","text":"text"}}`, string(m.Marshaler()))
}
