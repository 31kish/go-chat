package services

import (
	"go-chat/app/database"
	"go-chat/app/models"
	"log"
	"time"
)

// ChatEvent -
type ChatEvent struct {
	Base
}

const sendAtFormat = "2006/01/02 15:04"

// Create -
func (s ChatEvent) Create(e models.ChatEvent) (*models.ChatEvent, error) {
	db := *database.Connection

	result := db.Create(&e)

	if result.Error != nil {
		log.Printf("%#v", result.Error)
		return nil, result.Error
	}

	e.SendAt = e.CreatedAt.Format(sendAtFormat)
	return &e, nil
}

func (s ChatEvent) GetTodayEvent() ([]models.ChatEvent, error) {
	db := *database.Connection
	events := []models.ChatEvent{}

	to := time.Now()
	from := to.Add(-24 * time.Minute)

	query := db.Where("created_at BETWEEN ? AND ?", from, to).Order("created_at desc").Limit(15).Find(&events)
	count := query.RowsAffected

	if count == 0 {
		return nil, s.notFound()
	}

	events = reverse(events)
	events = addSendAt(events)

	return events, nil
}

func reverse(t []models.ChatEvent) []models.ChatEvent {
	r := make([]models.ChatEvent, len(t))
	index := len(t) - 1
	for i := 0; i < len(t); i++ {
		r[i] = t[index-i]
	}
	return r
}

func addSendAt(t []models.ChatEvent) []models.ChatEvent {
	r := make([]models.ChatEvent, len(t))
	for i, v := range t {
		r[i] = v
		r[i].SendAt = r[i].CreatedAt.Format(sendAtFormat)
	}
	return r
}
