package chatroom

import (
	"container/list"
	"go-chat/app/models"
	"go-chat/app/services"
	"log"
)

// Subscription -
type Subscription struct {
	Archive []models.ChatEvent      // All the events from the archive.
	New     <-chan models.ChatEvent // New events coming in.
}

// Cancel - Owner of a subscription must cancel it when they stop listening to events.
func (s Subscription) Cancel() {
	log.Println("Cancel()")
	unsubscribe <- s.New // Unsubscribe the channel.
	drain(s.New)         // Drain it, just in case there was a pending publish.
}

func newEvent(userID uint, typ int, msg string) models.ChatEvent {
	s := services.ChatEvent{}
	userService := services.User{}

	u, _ := userService.GetByID(int(userID))
	e, err := s.Create(models.ChatEvent{UserID: userID, UserName: u.Name, Type: typ, Text: msg})

	if err != nil {
		panic(err)
	}

	return *e
}

// Subscribe -
func Subscribe() Subscription {
	log.Println("Subscribe()")
	resp := make(chan Subscription)
	subscribe <- resp
	return <-resp
}

func Join(userID uint) {
	log.Println("Join")
	publish <- newEvent(userID, models.ChatEventType.Join(), "")
}

func Say(userID uint, message string) {
	publish <- newEvent(userID, models.ChatEventType.Message(), message)
}

func Leave(userID uint) {
	println("Leave")
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
			log.Println("subscribe")
			var events []models.ChatEvent
			for e := archive.Front(); e != nil; e = e.Next() {
				events = append(events, e.Value.(models.ChatEvent))
			}
			subscriber := make(chan models.ChatEvent, 10)
			subscribers.PushBack(subscriber)
			ch <- Subscription{events, subscriber}

		case event := <-publish:
			log.Println("publish")
			for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
				ch.Value.(chan models.ChatEvent) <- event
			}

			if archive.Len() >= archiveSize {
				archive.Remove(archive.Front())
			}

			archive.PushBack(event)

		case unsub := <-unsubscribe:
			log.Println("unsub")
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
