package controllers

import (
	"go-chat/app/chatroom"
	"strconv"

	"github.com/revel/revel"
	"golang.org/x/net/websocket"
)

// Chat -
type Chat struct {
	*revel.Controller
}

func init() {
	revel.InterceptMethod(Chat.before, revel.BEFORE)
}

func (c Chat) before() revel.Result {
	c.ViewArgs["title"] = "Chat Room"
	return c.Result
}

// Messages - Chat page.
func (c Chat) Messages() revel.Result {
	return c.Render()
}

// MessagesSocket -
func (c Chat) MessagesSocket(ws *websocket.Conn) revel.Result {
	// Make sure the websocket is valid.
	if ws == nil {
		return nil
	}

	// Join the room
	subscription := chatroom.Subscribe()
	defer subscription.Cancel()

	userId, _ := strconv.Atoi(c.Session["user_id"])
	id := uint(userId)
	chatroom.Join(id)
	defer chatroom.Leave(id)

	// Send down the archive.
	for _, event := range subscription.Archive {
		if websocket.JSON.Send(ws, &event) != nil {
			// They disconnected
			return nil
		}
	}

	// In rder to select between websocket messages and subscription events, we
	// need to stuff websocket events into a channel.
	newMessages := make(chan string)
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()

	// Now listen for new events from eigher the websocket or the chatroom
	for {
		select {
		case event := <-subscription.New:
			if websocket.JSON.Send(ws, &event) != nil {
				// They disconnected
				return nil
			}
		case msg, ok := <-newMessages:
			// If the channel is closed. they disconnected.
			if !ok {
				return nil
			}

			// Otherwise, say something.
			chatroom.Say(id, msg)
		}
	}
}
