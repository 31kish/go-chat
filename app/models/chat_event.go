package models

import "github.com/jinzhu/gorm"

// ChatEvent -
type ChatEvent struct {
	gorm.Model
	UserID   uint
	UserName string
	Type     int // "0:join", "1:leave", or "2:message"
	Text     string
	SendAt   string `gorm:"-"`
}

type chatEventTypeEnum int

const (
	join chatEventTypeEnum = iota
	leave
	message
)

type iChatEventType interface {
	Join() int
	Leave() int
	Message() int
}

// ChatEventType -
var ChatEventType iChatEventType = sChatEventType{}

type sChatEventType struct{}

func (t sChatEventType) Join() int {
	return int(join)
}

func (t sChatEventType) Leave() int {
	return int(leave)
}

func (t sChatEventType) Message() int {
	return int(message)
}
