package controllers

import (
	"go-chat/app/chatroom"
	"go-chat/app/routes"
	"go-chat/app/services"
	"log"
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
	if !c.loggedIn() {
		return c.Redirect(routes.App.Signin())
	}

	c.ViewArgs["title"] = "Chat Room"
	return c.Result
}

func (c Chat) loggedIn() bool {
	_, contains := c.Session["user_id"]
	return contains
}

// Messages - Chat page.
func (c Chat) Messages() revel.Result {
	log.Println("Chat Controller Messages")

	c.ViewArgs["user_id"] = c.Session["user_id"]
	c.ViewArgs["user_name"] = c.Session["user_name"]
	return c.Render()
}

// MessagesSocket -
func (c Chat) MessagesSocket(ws *websocket.Conn) revel.Result {
	log.Println("Chat Controller MessagesSocket")
	// Make sure the websocket is valid.
	if ws == nil {
		return nil
	}

	eventService := services.ChatEvent{}
	events, _ := eventService.GetTodayEvent()

	for _, event := range events {
		if websocket.JSON.Send(ws, &event) != nil {
			return nil
		}
	}

	// Join the room
	subscription := chatroom.Subscribe()
	defer subscription.Cancel()

	userID, _ := strconv.Atoi(c.Session["user_id"])
	id := uint(userID)
	chatroom.Join(id)
	defer chatroom.Leave(id)

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

	return nil
}
