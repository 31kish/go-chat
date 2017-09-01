package chatroom

import (
	"container/list"
	"go-chat/app/models"
)

// Subscription -
type Subscription struct {
	Archive []models.ChatEvent      // All the events from the archive.
	New     <-chan models.ChatEvent // New events coming in.
}

// Cancel - Owner of a subscription must cancel it when they stop listening to events.
func (s Subscription) Cancel() {
	unsubscribe <- s.New // Unsubscribe the channel.
	drain(s.New)         // Drain it, just in case there was a pending publish.
}

func newEvent(userID uint, typ int, msg string) models.ChatEvent {
	return models.ChatEvent{UserID: userID, Type: typ, Text: msg}
}

// Subscribe -
func Subscribe() Subscription {
	resp := make(chan Subscription)
	subscribe <- resp
	return <-resp
}

func Join(userID uint) {
	publish <- newEvent(userID, models.ChatEventType.Join(), "")
}

func Say(userID uint, message string) {
	publish <- newEvent(userID, models.ChatEventType.Message(), message)
}

func Leave(userID uint) {
	publish <- newEvent(userID, models.ChatEventType.Leave(), "")
}

const archiveSize = 10

var (
	// Send a channel here to get room events back. It will send the entire
	// archive initially, and then new messages as they come in.
	subscribe = make(chan (chan<- Subscription), 10)
	// Send a channel here to unsubscribe.
	unsubscribe = make(chan (<-chan models.ChatEvent), 10)
	// Send events here to publish them.
	publish = make(chan models.ChatEvent, 10)
)

// This function llops forever, handling the chat room pubsub
func chatroom() {
	archive := list.New()
	subscribers := list.New()

	for {
		select {
		case ch := <-subscribe:
			var events []models.ChatEvent
			for e := archive.Front(); e != nil; e = e.Next() {
				events = append(events, e.Value.(models.ChatEvent))
			}
			subscriber := make(chan models.ChatEvent, 10)
			subscribers.PushBack(subscriber)
			ch <- Subscription{events, subscriber}

		case event := <-publish:
			for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
				ch.Value.(chan models.ChatEvent) <- event
			}

			if archive.Len() >= archiveSize {
				archive.Remove(archive.Front())
			}

			archive.PushBack(event)

		case unsub := <-unsubscribe:
			for ch := subscribers.Front(); ch != nil; ch.Next() {
				if ch.Value.(chan models.ChatEvent) == unsub {
					subscribers.Remove(ch)
					break
				}
			}
		}
	}
}

func init() {
	go chatroom()
}

// Helpers

// Drains a given channel of any messages.
func drain(ch <-chan models.ChatEvent) {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
		default:
			return
		}
	}
}
