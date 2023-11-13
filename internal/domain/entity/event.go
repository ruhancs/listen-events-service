package entity

import (
	"github.com/asaskevich/govalidator"
	"github.com/rs/xid"
)

// Entidade de evento para coleta de eventos recebido de outros servicos
// Type = tipo do evento, Service = servico que gerou o evento, Date = data do evento, Status = informa se foi bem sucedido ou nao
type Event struct {
	ID      string `json:"id" valid:"required"`
	Type    string `json:"type" valid:"required"`
	Service string `json:"service" valid:"required"`
	Day     int    `json:"day" valid:"required"`
	Month   int    `json:"month" valid:"required"`
	Year    int    `json:"year" valid:"required"`
	Status  string `json:"status" valid:"required"`
}

func NewEvent(service, eventType, status string, day, month, year int) (*Event, error) {
	event := &Event{
		ID:      xid.New().String(),
		Type:    eventType,
		Service: service,
		Status:  status,
		Day:     day,
		Month:   month,
		Year:    year,
	}
	err := event.validate()
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (ev *Event) validate() error {
	_, err := govalidator.ValidateStruct(ev)
	if err != nil {
		return err
	}
	return nil
}
